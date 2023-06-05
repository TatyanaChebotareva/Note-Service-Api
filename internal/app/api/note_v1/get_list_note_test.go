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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetListNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		repoErrText = gofakeit.Phrase()
		repoErr     = errors.New(repoErrText)

		noteCnt = 3

		repoNotes  = make([]*model.Note, 0, noteCnt)
		validNotes = make([]*desc.Note, 0, noteCnt)
	)

	for i := 0; i < noteCnt; i++ {
		id := gofakeit.Int64()
		title := gofakeit.BeerName()
		text := gofakeit.BeerStyle()
		author := gofakeit.Name()
		createdAt := gofakeit.Date()
		updatedAt := gofakeit.Date()
		valid := true

		validNotes = append(validNotes, &desc.Note{
			Id: id,
			NoteInfo: &desc.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt),
		})

		repoNotes = append(repoNotes, &model.Note{
			Id: id,
			Info: &model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: valid,
			},
		})
	}

	validRes := &desc.GetListResponse{
		NoteList: validNotes,
	}

	noteMock := noteMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		noteMock.EXPECT().GetList(ctx).Return(repoNotes, nil),
		noteMock.EXPECT().GetList(ctx).Return(nil, repoErr),
	)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	t.Run("success case", func(t *testing.T) {
		// fmt.Println(req.GetId())
		res, err := api.GetList(ctx, &emptypb.Empty{})
		require.Equal(t, validRes, res)
		require.Nil(t, err)
	})

	t.Run("note repo err", func(t *testing.T) {
		_, err := api.GetList(ctx, &emptypb.Empty{})
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
