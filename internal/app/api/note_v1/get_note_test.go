package note_v1

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/model"
	noteMocks "github.com/TatyanaChebotareva/Note-Service-Api/internal/repository/note/mocks"
	note "github.com/TatyanaChebotareva/Note-Service-Api/internal/service"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id = gofakeit.Int64()

		title     = gofakeit.BeerName()
		text      = gofakeit.BeerStyle()
		author    = gofakeit.Name()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		tests = []struct {
			req      *desc.GetRequest
			validRes *desc.GetResponse
			repoRes  *model.Note
		}{
			{
				req: &desc.GetRequest{
					Id: id,
				},

				validRes: &desc.GetResponse{
					Note: &desc.Note{
						Id: id,
						NoteInfo: &desc.NoteInfo{
							Title:  title,
							Text:   text,
							Author: author,
						},
						CreatedAt: timestamppb.New(createdAt),
						UpdatedAt: timestamppb.New(updatedAt),
					},
				},

				repoRes: &model.Note{
					Id: id,
					Info: &model.NoteInfo{
						Title:  title,
						Text:   text,
						Author: author,
					},
					CreatedAt: createdAt,
					UpdatedAt: sql.NullTime{
						Time:  updatedAt,
						Valid: true,
					},
				},
			},
		}

		repoErrText = gofakeit.Phrase()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		for _, tc := range tests {
			noteMock.EXPECT().Get(ctx, id).Return(tc.repoRes, nil)
			res, err := api.Get(ctx, tc.req)
			require.Equal(t, tc.validRes, res)
			require.Nil(t, err)
		}
	})

	t.Run("note repo err", func(t *testing.T) {
		for _, tc := range tests {
			noteMock.EXPECT().Get(ctx, id).Return(nil, repoErr)
			_, err := api.Get(ctx, tc.req)
			require.NotNil(t, err)
			require.Equal(t, repoErrText, err.Error())
		}
	})
}
