package model

import (
	"database/sql"
	"time"
)

type NoteInfo struct {
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
}

type UpdateNoteInfo struct {
	Id     int64          `db:"id"`
	Title  sql.NullString `db:"title"`
	Text   sql.NullString `db:"text"`
	Author sql.NullString `db:"author"`
}

type Note struct {
	Id        int64        `db:"id"`
	Info      *NoteInfo    `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
