package note

import (
	"context"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/converter"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (s *Service) GetList(ctx context.Context) (*desc.GetListResponse, error) {
	notes, err := s.noteRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}

	descNotes := make([]*desc.Note, 0, len(notes))

	for _, note := range notes {
		descNotes = append(descNotes, converter.ToDescNote(note))
	}

	return &desc.GetListResponse{
		NoteList: descNotes,
	}, nil
}
