package note_v1

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println("Update")

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Update(noteTable).PlaceholderFormat(sq.Dollar).
		Set("title", request.Note.GetTitle()).
		Set("text", request.Note.GetText()).
		Set("author", request.Note.GetAuthor()).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": request.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = db.DB.Exec(query, args...)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
