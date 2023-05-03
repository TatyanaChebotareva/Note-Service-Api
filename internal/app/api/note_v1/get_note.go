package note_v1

import (
	"context"
	"fmt"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("GetNote")
	fmt.Println("Id: ", req.GetId())

	return &desc.GetNoteResponse{
		Title:  "Doctors visit",
		Text:   "15.05.2023",
		Author: "Aibolyt",
	}, nil
}
