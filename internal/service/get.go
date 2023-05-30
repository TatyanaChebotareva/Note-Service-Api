package note

import (
	"context"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/converter"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (s *Service) Get(ctx context.Context, id int64) (*desc.GetResponse, error) {
	note, err := s.noteRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: converter.ToDescNote(note),
	}, nil
}
