package note_v1

import (
	"context"
	"fmt"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/converter"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println("Update")

	err := n.noteService.Update(ctx, converter.ToUpdateNoteInfo(req.GetNote()))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
