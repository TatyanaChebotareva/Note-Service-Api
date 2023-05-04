package note_v1

import (
	"context"
	"fmt"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println("Update")
	fmt.Println("Id:", request.GetId())
	fmt.Println("Title:", request.Note.GetTitle())
	fmt.Println("Text:", request.Note.GetText())
	fmt.Println("Author:", request.Note.GetAuthor())

	return &emptypb.Empty{}, nil
}
