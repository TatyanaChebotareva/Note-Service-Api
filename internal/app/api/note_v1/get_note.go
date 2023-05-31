package note_v1

import (
	"context"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/converter"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	res, err := n.noteService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: converter.ToDescNote(res),
	}, nil
}
