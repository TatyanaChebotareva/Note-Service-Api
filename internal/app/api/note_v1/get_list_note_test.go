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

type testCase struct {
	testName  string
	validRes  *desc.GetListResponse
	repoNotes []*model.Note
	error     error
}

func TestGetListNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		repoErrText = gofakeit.Phrase()

		noteCnt = 3

		repoNotes  = make([]*model.Note, 0, noteCnt)
		validNotes = make([]*desc.Note, 0, noteCnt)

		tests []testCase
	)

	// cycle for success test data generation
	for i := 0; i < noteCnt; i++ {
		id := gofakeit.Int64()
		title := gofakeit.BeerName()
		text := gofakeit.BeerStyle()
		author := gofakeit.Name()
		createdAt := gofakeit.Date()
		updatedAt := gofakeit.Date()

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
				Valid: true,
			},
		})
	}

	validRes := &desc.GetListResponse{
		NoteList: validNotes,
	}

	tests = append(tests, testCase{
		testName:  "success case",
		validRes:  validRes,
		repoNotes: repoNotes,
		error:     nil,
	})

	tests = append(tests, testCase{
		testName:  "failed case",
		validRes:  nil,
		repoNotes: nil,
		error:     errors.New(repoErrText),
	})

	noteMock := noteMocks.NewMockRepository(mockCtrl)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			noteMock.EXPECT().GetList(ctx).Return(tc.repoNotes, tc.error)
			res, err := api.GetList(ctx, &emptypb.Empty{})
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
