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

		title  = gofakeit.BeerName()
		text   = gofakeit.BeerStyle()
		author = gofakeit.Name()

		id = gofakeit.Int64()

		repoErrText = gofakeit.Phrase()

		tests = []struct {
			testName string
			req      *desc.CreateRequest
			repoReq  *model.NoteInfo
			validRes *desc.CreateResponse
			repoRes  int64
			error    error
		}{
			{
				testName: "success case",
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
				repoRes: id,
				error:   nil,
			},
			{
				testName: "failed case",
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
				validRes: nil,
				repoRes:  int64(0),
				error:    errors.New(repoErrText),
			},
		}
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			noteMock.EXPECT().Create(ctx, tc.repoReq).Return(tc.repoRes, tc.error)
			res, err := api.Create(ctx, tc.req)
			if tc.error == nil {
				require.Equal(t, tc.validRes, res)
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.Equal(t, repoErrText, err.Error())
			}
		})
	}
}
