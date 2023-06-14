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

		repoErrText = gofakeit.Phrase()

		tests = []struct {
			testName string
			req      *desc.GetRequest
			validRes *desc.GetResponse
			repoRes  *model.Note
			error    error
		}{
			{
				testName: "success case",
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
				error: nil,
			},
			{
				testName: "failed case",
				req: &desc.GetRequest{
					Id: id,
				},
				validRes: nil,
				repoRes:  nil,
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
			noteMock.EXPECT().Get(ctx, id).Return(tc.repoRes, tc.error)
			res, err := api.Get(ctx, tc.req)
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
