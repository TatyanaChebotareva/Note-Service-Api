package note_v1

import (
	"context"
	"fmt"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get")
	fmt.Println("Id: ", req.GetId())

	note := desc.Note{
		Title:  "Doctors visit",
		Text:   "15.05.2023",
		Author: "Aibolyt",
	}

	return &desc.GetResponse{
		Note: &note,
	}, nil
}
