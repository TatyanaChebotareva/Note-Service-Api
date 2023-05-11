package note

import "github.com/TatyanaChebotareva/Note-Service-Api/internal/repository"

type Service struct {
	noteRepository repository.NoteRepository
}

func NewService(noteRepository repository.NoteRepository) *Service {
	return &Service{
		noteRepository: noteRepository,
	}
}
