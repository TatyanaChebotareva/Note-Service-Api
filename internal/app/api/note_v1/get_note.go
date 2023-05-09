package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func (n *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	builder := sq.Select("title, text, author").From(noteTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	note := desc.Note{}

	row.Next()
	err = row.Scan(&note.Title, &note.Text, &note.Author)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Note: &note,
	}, nil
}
