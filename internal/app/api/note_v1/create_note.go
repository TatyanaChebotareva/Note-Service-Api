package note_v1

import (
	"context"
	"fmt"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Println("Create")
	fmt.Println("title:", req.GetTitle())
	fmt.Println("text:", req.GetText())
	fmt.Println("author:", req.GetAuthor())

	return &desc.CreateResponse{
		Id: 1,
	}, nil
}
