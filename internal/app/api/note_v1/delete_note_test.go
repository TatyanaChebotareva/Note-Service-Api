package note_v1

import (
	"context"
	"errors"
	"fmt"
	"testing"

	noteMocks "github.com/TatyanaChebotareva/Note-Service-Api/internal/repository/note/mocks"
	note "github.com/TatyanaChebotareva/Note-Service-Api/internal/service"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id = gofakeit.Int64()

		req = &desc.DeleteRequest{
			Id: id,
		}

		repoErrText = gofakeit.Phrase()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().Delete(ctx, id).Return(nil),
		noteMock.EXPECT().Delete(ctx, id).Return(repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		fmt.Println(req.GetId())
		_, err := api.Delete(ctx, req)
		require.Nil(t, err)
	})

	t.Run("note repo err", func(t *testing.T) {
		_, err := api.Delete(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}