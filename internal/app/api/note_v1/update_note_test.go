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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdateNote(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id = gofakeit.Int64()

		tests = map[string]struct {
			title  string
			text   string
			author string
		}{
			"correct data": {gofakeit.BeerName(), gofakeit.BeerStyle(), gofakeit.Name()},
			"one nullable": {"", gofakeit.BeerStyle(), gofakeit.Name()},
			"all nullable": {"", "", ""},
		}

		repoErrText = gofakeit.Phrase()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)

	api := newMockNoteV1(Note{
		noteService: note.NewMockNoteService(noteMock),
	})

	for name, tc := range tests {
		req := &desc.UpdateRequest{
			Note: &desc.UpdateNoteInfo{
				Id: id,
			},
		}

		repoReq := &model.UpdateNoteInfo{
			Id: id,
		}

		if tc.title != "" {
			req.GetNote().Title = &wrapperspb.StringValue{Value: tc.title}
			repoReq.Title.String = tc.title
			repoReq.Title.Valid = true
		}

		if tc.text != "" {
			req.GetNote().Text = &wrapperspb.StringValue{Value: tc.text}
			repoReq.Text.String = tc.text
			repoReq.Text.Valid = true
		}

		if tc.author != "" {
			req.GetNote().Author = &wrapperspb.StringValue{Value: tc.author}
			repoReq.Author.String = tc.author
			repoReq.Author.Valid = true
		}

		t.Run("success case "+name, func(t *testing.T) {
			noteMock.EXPECT().Update(ctx, repoReq).Return(nil)
			_, err := api.Update(ctx, req)
			require.Nil(t, err)
		})

		t.Run("note repo err "+name, func(t *testing.T) {
			noteMock.EXPECT().Update(ctx, repoReq).Return(repoErr)
			_, err := api.Update(ctx, req)
			require.NotNil(t, err)
			require.Equal(t, repoErrText, err.Error())
		})
	}
}
