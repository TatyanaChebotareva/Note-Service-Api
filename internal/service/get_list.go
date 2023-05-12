package note

import (
	"context"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (s *Service) GetList(ctx context.Context) (*desc.GetListResponse, error) {
	notes, timestamps, err := s.noteRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	return &desc.GetListResponse{
		NoteList:      notes,
		TimestampList: timestamps,
	}, nil
}
