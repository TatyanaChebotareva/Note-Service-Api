package note

import (
	"context"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (s *Service) Update(ctx context.Context, req *desc.UpdateRequest) error {
	err := s.noteRepository.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
