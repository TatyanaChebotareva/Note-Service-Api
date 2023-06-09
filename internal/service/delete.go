package note

import (
	"context"
)

func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.noteRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
