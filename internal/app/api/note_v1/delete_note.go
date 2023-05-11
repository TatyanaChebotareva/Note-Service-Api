package note_v1

import (
	"context"
	"fmt"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	fmt.Println("Delete")

	res, err := n.noteService.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
