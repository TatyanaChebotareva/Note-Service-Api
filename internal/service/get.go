package note

import (
	"context"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (s *Service) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	note, err := s.noteRepository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: note,
	}, nil
}
