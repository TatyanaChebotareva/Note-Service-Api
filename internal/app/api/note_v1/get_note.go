package note_v1

import (
	"context"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	res, err := n.noteService.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
