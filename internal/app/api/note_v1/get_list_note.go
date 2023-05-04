package note_v1

import (
	"context"
	"fmt"

	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetList(ctx context.Context, in *emptypb.Empty) (*desc.GetListResponse, error) {
	fmt.Println("GetList")

	noteList := []*desc.Note{
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

	return &desc.GetListResponse{
		NoteList: noteList,
	}, nil
}
