package converter

import (
	"database/sql"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/model"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToNoteInfo(noteInfo *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
	}
}

func ToDescNoteInfo(noteInfo *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
	}
}

func ToNote(note *desc.Note) *model.Note {
	var updatedAt sql.NullTime

	if note.GetUpdatedAt().IsValid() {
		updatedAt.Time = note.GetUpdatedAt().AsTime()
	}

	return &model.Note{
		Id:        note.GetId(),
		Info:      ToNoteInfo(note.GetNoteInfo()),
		CreatedAt: note.GetCreatedAt().AsTime(),
		UpdatedAt: updatedAt,
	}
}

func ToDescNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.Id,
		NoteInfo:  ToDescNoteInfo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUpdateNoteInfo(updateNoteInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	var title, text, author sql.NullString

	if updateNoteInfo.GetTitle() != nil {
		title.String = updateNoteInfo.GetTitle().Value
		title.Valid = true
	}

	if updateNoteInfo.GetText() != nil {
		text.String = updateNoteInfo.GetText().Value
		text.Valid = true
	}

	if updateNoteInfo.GetAuthor() != nil {
		author.String = updateNoteInfo.GetAuthor().Value
		author.Valid = true
	}

	return &model.UpdateNoteInfo{
		Id:     updateNoteInfo.GetId(),
		Title:  title,
		Text:   text,
		Author: author,
	}
}

func ToDescUpdateNoteInfo(noteInfo *model.UpdateNoteInfo) *desc.UpdateNoteInfo {
	var title, text, author wrapperspb.StringValue

	if noteInfo.Title.Valid {
		title.Value = noteInfo.Title.String
	}

	if noteInfo.Text.Valid {
		text.Value = noteInfo.Text.String
	}

	if noteInfo.Author.Valid {
		author.Value = noteInfo.Author.String
	}

	return &desc.UpdateNoteInfo{
		Id:     noteInfo.Id,
		Title:  &title,
		Text:   &text,
		Author: &author,
	}
}
