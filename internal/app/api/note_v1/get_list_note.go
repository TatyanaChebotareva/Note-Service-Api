package note_v1

import (
	"context"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/converter"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetList(ctx context.Context, in *emptypb.Empty) (*desc.GetListResponse, error) {
	res, err := n.noteService.GetList(ctx)
	if err != nil {
		return nil, err
	}

	descNotes := make([]*desc.Note, 0, len(res))

	for _, note := range res {
		descNotes = append(descNotes, converter.ToDescNote(note))
	}

	return &desc.GetListResponse{
		NoteList: descNotes,
	}, nil
}
