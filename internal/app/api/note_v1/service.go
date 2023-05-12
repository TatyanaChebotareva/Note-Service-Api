package note_v1

import (
	note "github.com/TatyanaChebotareva/Note-Service-Api/internal/service"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

type Note struct {
	desc.UnimplementedNoteV1Server
	noteService *note.Service
}

func NewNote(noteService *note.Service) *Note {
	return &Note{
		noteService: noteService,
	}
}
