package note_v1

import (
	"context"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	res, err := n.noteService.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
