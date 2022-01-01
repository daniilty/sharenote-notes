package core

import "github.com/daniilty/sharenote-notes/internal/mongo"

func (n *Note) toDB() *mongo.Note {
	return &mongo.Note{
		ID:        n.ID,
		Data:      n.Data,
		Name:      n.Name,
		Timestamp: n.Timestamp,
	}
}

func convertDBNoteToService(note *mongo.Note) *Note {
	return &Note{
		ID:        note.ID,
		Data:      note.Data,
		Name:      note.Name,
		Timestamp: note.Timestamp,
	}
}
