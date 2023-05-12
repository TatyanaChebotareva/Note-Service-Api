package note

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	Note = "note"
)

type Repository interface {
	Create(ctx context.Context, req *desc.CreateRequest) (int64, error)
	Get(ctx context.Context, req *desc.GetRequest) (*desc.Note, *desc.Timestamp, error)
	GetList(ctx context.Context) ([]*desc.Note, []*desc.Timestamp, error)
	Delete(ctx context.Context, req *desc.DeleteRequest) error
	Update(ctx context.Context, req *desc.UpdateRequest) error
}

type repository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(ctx context.Context, req *desc.CreateRequest) (int64, error) {
	builder := sq.Insert(Note).
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

func (r *repository) Get(ctx context.Context, req *desc.GetRequest) (*desc.Note, *desc.Timestamp, error) {
	builder := sq.Select("title, text, author, created_at, updated_at").From(Note).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer row.Close()

	note := desc.Note{}
	timestamp := desc.Timestamp{}
	var createTime time.Time
	var updateTime sql.NullTime

	row.Next()
	err = row.Scan(&note.Title, &note.Text, &note.Author, &createTime, &updateTime)
	if err != nil {
		return nil, nil, err
	}

	timestamp.CreatedAt = timestamppb.New(createTime)
	if updateTime.Valid {
		timestamp.UpdatedAt = timestamppb.New(updateTime.Time)
	}

	return &note, &timestamp, nil
}

func (r *repository) GetList(ctx context.Context) ([]*desc.Note, []*desc.Timestamp, error) {
	builder := sq.Select("title, text, author, created_at, updated_at").From(Note)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, nil, err
	}

	row, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer row.Close()

	var noteList []*desc.Note
	var timestampList []*desc.Timestamp

	var createTime time.Time
	var updateTime sql.NullTime

	for row.Next() {
		note := new(desc.Note)
		timestamp := new(desc.Timestamp)

		err = row.Scan(&note.Title, &note.Text, &note.Author, &createTime, &updateTime)
		if err != nil {
			return nil, nil, err
		}

		timestamp.CreatedAt = timestamppb.New(createTime)
		if updateTime.Valid {
			timestamp.UpdatedAt = timestamppb.New(updateTime.Time)
		}

		noteList = append(noteList, note)
		timestampList = append(timestampList, timestamp)
	}

	return noteList, timestampList, nil
}

func (r *repository) Delete(ctx context.Context, req *desc.DeleteRequest) error {
	builder := sq.Delete(Note).PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.DB.ExecContext(ctx, query, args...); err != nil {
		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, req *desc.UpdateRequest) error {
	builder := sq.Update(Note).PlaceholderFormat(sq.Dollar).
		Set("title", req.Note.GetTitle()).
		Set("text", req.Note.GetText()).
		Set("author", req.Note.GetAuthor()).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": req.GetId()})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.DB.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}
