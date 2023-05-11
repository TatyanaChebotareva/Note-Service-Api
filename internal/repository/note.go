package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/TatyanaChebotareva/Note-Service-Api/internal/repository/table"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
)

type NoteRepository interface {
	Create(ctx context.Context, req *desc.CreateRequest) (int64, error)
	Get(ctx context.Context, req *desc.GetRequest) (*desc.Note, error)
	GetList(ctx context.Context) ([]*desc.Note, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) error
	Update(ctx context.Context, req *desc.UpdateRequest) error
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, req *desc.CreateRequest) (int64, error) {
	builder := sq.Insert(table.Note).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.Note.GetTitle(), req.Note.GetText(), req.Note.GetAuthor()).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Get(ctx context.Context, req *desc.GetRequest) (*desc.Note, error) {
	builder := sq.Select("title, text, author").From(table.Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
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

	return &note, nil
}

func (r *repository) GetList(ctx context.Context) ([]*desc.Note, error) {
	builder := sq.Select("title, text, author").From(table.Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var noteList []*desc.Note

	for row.Next() {
		note := new(desc.Note)

		err = row.Scan(&note.Title, &note.Text, &note.Author)
		if err != nil {
			return nil, err
		}

		noteList = append(noteList, note)
	}

	return noteList, nil
}

func (r *repository) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	builder := sq.Delete(table.Note).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.DB.Exec(query, args...); err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, req *desc.UpdateRequest) error {
	builder := sq.Update(table.Note).PlaceholderFormat(sq.Dollar).
		Set("title", req.Note.GetTitle()).
		Set("text", req.Note.GetText()).
		Set("author", req.Note.GetAuthor()).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.DB.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}
