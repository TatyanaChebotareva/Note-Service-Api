package note_v1

import (
	"context"
	"fmt"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
)

func (n *Note) GetListNote(cts context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {

	fmt.Println("GetListNote")

	noteList := []*desc.GetNoteResponse{
		{
			Title:  "Wow!",
			Text:   "It's working",
			Author: "Tanya",
		},
		{
			Title:  "Doctors visit",
			Text:   "15.05.2023",
			Author: "Aibolyt",
		},
		{
			Title:  "Don't forget",
			Text:   "Better late than never",
			Author: "Tanya",
		},
	}

	return &desc.GetListNoteResponse{
		NoteList: noteList,
	}, nil
}
