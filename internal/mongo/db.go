package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ DB = (*DBImpl)(nil)

type DB interface {
	// AddNote - add note to database.
	AddNote(context.Context, *Note) error
	// GetNote - get note by id.
	GetNote(context.Context, string) (*Note, bool, error)
	// DeleteNote - delete note by id.
	DeleteNote(context.Context, string) (bool, error)
	// UpdateNote - update note by id.
	UpdateNote(context.Context, *Note) (bool, error)
}

type DBImpl struct {
	mongoDB         *mongo.Database
	notesCollection *mongo.Collection
}

func NewDBImpl(db *mongo.Database, notesCollection *mongo.Collection) *DBImpl {
	return &DBImpl{
		mongoDB:         db,
		notesCollection: notesCollection,
	}
}

func Connect(ctx context.Context, addr string) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(addr))
}
