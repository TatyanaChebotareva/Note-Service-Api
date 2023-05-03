package note_v1

import (
	"context"
	"fmt"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) UpdateNote(ctx context.Context, request *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {

	fmt.Println("UpdateNote")
	fmt.Println("Id:", request.GetId())
	fmt.Println("Text:", request.GetText())

	return &desc.UpdateNoteResponse{
		Result: "Note was successfully updated",
	}, nil
}
