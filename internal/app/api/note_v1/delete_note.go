package note_v1

import (
	"context"
	"fmt"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Println("Delete")
	fmt.Println("Id: ", req.GetId())

	return &emptypb.Empty{}, nil
}
