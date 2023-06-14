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

func NewMockNoteService(deps ...interface{}) *Service {
	is := Service{}

	for _, v := range deps {
		switch s := v.(type) {
		case note.Repository:
			is.noteRepository = s
		}
	}

	return &is
}
