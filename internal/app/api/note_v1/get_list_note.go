package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetList(ctx context.Context, in *emptypb.Empty) (*desc.GetListResponse, error) {
	fmt.Println("GetList")

	var noteList []*desc.Note
	var note *desc.Note

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("title, text, author").From(noteTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		note = new(desc.Note)

		err = row.Scan(&note.Title, &note.Text, &note.Author)
		if err != nil {
			return nil, err
		}

		noteList = append(noteList, note)
	}

	return &desc.GetListResponse{
		NoteList: noteList,
	}, nil
}
