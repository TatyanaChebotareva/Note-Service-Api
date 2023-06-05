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

		valid = true

		title  = gofakeit.BeerName()
		text   = gofakeit.BeerStyle()
		author = gofakeit.Name()

		req = &desc.UpdateRequest{
			Note: &desc.UpdateNoteInfo{
				Id: id,
				Title: &wrapperspb.StringValue{
					Value: title,
				},
				Text: &wrapperspb.StringValue{
					Value: text,
				},
				Author: &wrapperspb.StringValue{
					Value: author,
				},
			},
		}

		repoReq = &model.UpdateNoteInfo{
			Id: id,
			Title: sql.NullString{
				String: title,
				Valid:  valid,
			},
			Text: sql.NullString{
				String: text,
				Valid:  valid,
			},
			Author: sql.NullString{
				String: author,
				Valid:  valid,
			},
		}

		repoErrText = gofakeit.Phrase()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().Update(ctx, repoReq).Return(nil),
		noteMock.EXPECT().Update(ctx, repoReq).Return(repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		// fmt.Println(req.GetId())
		_, err := api.Update(ctx, req)
		require.Nil(t, err)
	})

	t.Run("note repo err", func(t *testing.T) {
		_, err := api.Update(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
