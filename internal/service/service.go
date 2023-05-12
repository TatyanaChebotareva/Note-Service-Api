package note

import (
	"github.com/TatyanaChebotareva/Note-Service-Api/internal/repository/note"
)

type Service struct {
	noteRepository note.Repository
}

func NewService(noteRepository note.Repository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
