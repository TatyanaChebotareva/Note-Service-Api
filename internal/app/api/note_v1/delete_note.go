package note_v1

import (
	"context"
	"fmt"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {

	fmt.Println("DeleteNote")
	fmt.Println("Id: ", req.GetId())

	return &desc.DeleteNoteResponse{
		Result: "Note was successfully deleted",
	}, nil
}
