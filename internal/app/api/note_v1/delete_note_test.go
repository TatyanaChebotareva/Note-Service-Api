package note_v1

import (
	"context"
	"errors"
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

		id          = gofakeit.Int64()
		repoErrText = gofakeit.Phrase()

		tests = []struct {
			testName string
			req      *desc.DeleteRequest
			error    error
		}{
			{
				testName: "success case",
				req: &desc.DeleteRequest{
					Id: id,
				},
				error: nil,
			},
			{
				testName: "failed case",
				req: &desc.DeleteRequest{
					Id: id,
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
			noteMock.EXPECT().Delete(ctx, id).Return(tc.error)
			_, err := api.Delete(ctx, tc.req)
			if tc.error == nil {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.Equal(t, repoErrText, err.Error())
			}
		})
	}
}
