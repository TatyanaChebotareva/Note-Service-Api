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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id = gofakeit.Int64()

		title       = gofakeit.BeerName()
		text        = gofakeit.BeerStyle()
		author      = gofakeit.Name()
		repoErrText = gofakeit.Phrase()

		tests = []struct {
			testName string
			req      *desc.UpdateRequest
			repoReq  *model.UpdateNoteInfo
			error    error
		}{
			{
				testName: "success case correct data",
				req: &desc.UpdateRequest{
					Note: &desc.UpdateNoteInfo{
						Id:     id,
						Title:  &wrapperspb.StringValue{Value: title},
						Text:   &wrapperspb.StringValue{Value: text},
						Author: &wrapperspb.StringValue{Value: author},
					},
				},
				repoReq: &model.UpdateNoteInfo{
					Id: id,
					Title: sql.NullString{
						String: title,
						Valid:  true,
					},
					Text: sql.NullString{
						String: text,
						Valid:  true,
					},
					Author: sql.NullString{
						String: author,
						Valid:  true,
					},
				},
				error: nil,
			},
			{
				testName: "success case one nullable",
				req: &desc.UpdateRequest{
					Note: &desc.UpdateNoteInfo{
						Id:     id,
						Title:  nil,
						Text:   &wrapperspb.StringValue{Value: text},
						Author: &wrapperspb.StringValue{Value: author},
					},
				},
				repoReq: &model.UpdateNoteInfo{
					Id: id,
					Title: sql.NullString{
						String: "",
						Valid:  false,
					},
					Text: sql.NullString{
						String: text,
						Valid:  true,
					},
					Author: sql.NullString{
						String: author,
						Valid:  true,
					},
				},
				error: nil,
			},
			{
				testName: "success case all nullable",
				req: &desc.UpdateRequest{
					Note: &desc.UpdateNoteInfo{
						Id:     id,
						Title:  nil,
						Text:   nil,
						Author: nil,
					},
				},
				repoReq: &model.UpdateNoteInfo{
					Id: id,
					Title: sql.NullString{
						String: "",
						Valid:  false,
					},
					Text: sql.NullString{
						String: "",
						Valid:  false,
					},
					Author: sql.NullString{
						String: "",
						Valid:  false,
					},
				},
				error: nil,
			},
			{
				testName: "failed case",
				req: &desc.UpdateRequest{
					Note: &desc.UpdateNoteInfo{
						Id:     id,
						Title:  nil,
						Text:   nil,
						Author: nil,
					},
				},
				repoReq: &model.UpdateNoteInfo{
					Id: id,
					Title: sql.NullString{
						String: "",
						Valid:  false,
					},
					Text: sql.NullString{
						String: "",
						Valid:  false,
					},
					Author: sql.NullString{
						String: "",
						Valid:  false,
					},
				},
				error: errors.New(repoErrText),
			},
		}
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			noteMock.EXPECT().Update(ctx, tc.repoReq).Return(tc.error)
			_, err := api.Update(ctx, tc.req)
			if tc.error == nil {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.Equal(t, repoErrText, err.Error())
			}
		})
	}
}
