package server

import (
	schema "github.com/daniilty/sharenote-grpc-schema"
	"github.com/daniilty/sharenote-notes/internal/core"
)

func convertCoreNoteToPB(note *core.Note) *schema.Note {
	return &schema.Note{
		Id:        note.ID,
		Data:      note.Data,
		Name:      note.Name,
		Timestamp: note.Timestamp,
	}
}

func convertPBNoteToCore(note *schema.Note) *core.Note {
	return &core.Note{
		ID:        note.Id,
		Data:      note.Data,
		Name:      note.Name,
		Timestamp: note.Timestamp,
	}
}

func convertCoreNotesToPB(notes []*core.Note) []*schema.Note {
	converted := make([]*schema.Note, 0, len(notes))

	for i := range notes {
		converted = append(converted, convertCoreNoteToPB(notes[i]))
	}

	return converted
}
