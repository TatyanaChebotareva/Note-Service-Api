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

		req = &desc.CreateRequest{
			Note: &desc.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		}

		repoReq = &model.NoteInfo{
			Title:  title,
			Text:   text,
			Author: author,
		}

		validRes = &desc.CreateResponse{
			Id: id,
		}

		repoErr = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().Create(ctx, repoReq).Return(id, nil),
		noteMock.EXPECT().Create(ctx, repoReq).Return(int64(0), repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		// fmt.Println(req.GetNote().GetTitle(), ";", req.GetNote().GetText(), ";", req.GetNote().GetAuthor())
		res, err := api.Create(ctx, req)
		require.Equal(t, validRes, res)
		require.Nil(t, err)
	})

	t.Run("note repo err", func(t *testing.T) {
		_, err := api.Create(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
