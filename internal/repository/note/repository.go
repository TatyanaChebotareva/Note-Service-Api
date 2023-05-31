package note

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/TatyanaChebotareva/Note-Service-Api/internal/model"
	"github.com/TatyanaChebotareva/Note-Service-Api/internal/pkg/db"
)

const (
	tableName = "note"
)

type Repository interface {
	Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	GetList(ctx context.Context) ([]*model.Note, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, req *model.UpdateNoteInfo) error
}

type repository struct {
	client db.Client
}

func NewNoteRepository(client db.Client) Repository {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, noteInfo *model.NoteInfo) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(noteInfo.Title, noteInfo.Text, noteInfo.Author).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "Create_note",
		QueryRaw: query,
	}

	row, err := r.client.DB().QueryContext(ctx, q, args...)
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

func (r *repository) Get(ctx context.Context, id int64) (*model.Note, error) {
	builder := sq.Select("id, title, text, author, created_at, updated_at").From(tableName).
		PlaceholderFormat(sq.Dollar).Where(sq.Eq{"id": id}).Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "Get_note",
		QueryRaw: query,
	}

	var note model.Note

	err = r.client.DB().GetContext(ctx, &note, q, args...)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *repository) GetList(ctx context.Context) ([]*model.Note, error) {
	builder := sq.Select("id, title, text, author, created_at, updated_at").From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetList_note",
		QueryRaw: query,
	}

	var notes []*model.Note

	err = r.client.DB().SelectContext(ctx, &notes, q, args...)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Delete_note",
		QueryRaw: query,
	}

	if _, err := r.client.DB().ExecContext(ctx, q, args...); err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, req *model.UpdateNoteInfo) error {
	builder := sq.Update(tableName).PlaceholderFormat(sq.Dollar).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": req.Id})

	if req.Title.Valid {
		builder = builder.Set("title", req.Title.String)
	}

	if req.Text.Valid {
		builder = builder.Set("text", req.Text.String)
	}

	if req.Author.Valid {
		builder = builder.Set("author", req.Author.String)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "Update_note",
		QueryRaw: query,
	}

	_, err = r.client.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return err
	}

	return nil
}
