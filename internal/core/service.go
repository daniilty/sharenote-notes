package core

import (
	"context"

	"github.com/daniilty/sharenote-notes/internal/mongo"
)

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	// AddNote - add note to database.
	AddNote(context.Context, *Note) error
	// GetNote - get note by id.
	GetNote(context.Context, string) (*Note, bool, error)
	// GetNotes - get notes by id.
	GetNotes(context.Context, []string) ([]*Note, error)
	// DeleteNote - delete note by id.
	DeleteNote(context.Context, string) (bool, error)
	// DeleteNotes - delete notes by id.
	DeleteNotes(context.Context, []string) error
	// UpdateNote - update note by id.
	UpdateNote(context.Context, *Note) (bool, error)
}

type ServiceImpl struct {
	db mongo.DB
}

func NewServiceImpl(db mongo.DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
