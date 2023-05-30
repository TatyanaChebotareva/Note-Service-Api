package note

import (
	"context"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/model"
)

func (s *Service) Update(ctx context.Context, noteInfo *model.UpdateNoteInfo) error {
	err := s.noteRepository.Update(ctx, noteInfo)
	if err != nil {
		return err
	}

	return nil
}
