package note_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/model"
	noteMocks "github.com/TatyanaChebotareva/Note-Service-Api/internal/repository/note/mocks"
	note "github.com/TatyanaChebotareva/Note-Service-Api/internal/service"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id = gofakeit.Int64()

		title  = gofakeit.BeerName()
		text   = gofakeit.BeerStyle()
		author = gofakeit.Name()

		repoErrText = gofakeit.Phrase()
		repoErr     = errors.New(repoErrText)

		tests = []struct {
			req      *desc.CreateRequest
			repoReq  *model.NoteInfo
			validRes *desc.CreateResponse
		}{
			{
				req: &desc.CreateRequest{
					Note: &desc.NoteInfo{
						Title:  title,
						Text:   text,
						Author: author,
					},
				},
				repoReq: &model.NoteInfo{
					Title:  title,
					Text:   text,
					Author: author,
				},
				validRes: &desc.CreateResponse{
					Id: id,
				},
			},
		}
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		for _, tc := range tests {
			noteMock.EXPECT().Create(ctx, tc.repoReq).Return(id, nil)
			res, err := api.Create(ctx, tc.req)
			require.Equal(t, tc.validRes, res)
			require.Nil(t, err)
		}
	})

	t.Run("note repo err", func(t *testing.T) {
		for _, tc := range tests {
			noteMock.EXPECT().Create(ctx, tc.repoReq).Return(int64(0), repoErr)
			_, err := api.Create(ctx, tc.req)
			require.NotNil(t, err)
			require.Equal(t, repoErrText, err.Error())
		}
	})
}
