package mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	ID        string `bson:"_id"`
	Name      string `bson:"name"`
	Data      string `bson:"data"`
	Timestamp int64  `bson:"ts"`
}

func (n *Note) toBSOND() bson.D {
	return bson.D{
		{Key: "name", Value: n.Name},
		{Key: "data", Value: n.Data},
		{Key: "ts", Value: n.Timestamp},
	}
}

func (d *DBImpl) GetNote(ctx context.Context, id string) (*Note, bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, true, fmt.Errorf("bad object id: %s", id)
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	res := d.notesCollection.FindOne(ctx, filter)
	note := &Note{}

	err = res.Decode(note)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return note, true, err
		}

		return nil, false, err
	}

	return note, true, nil
}

func (d *DBImpl) AddNote(ctx context.Context, note *Note) error {
	_, err := d.notesCollection.InsertOne(ctx, note.toBSOND())

	return err
}

func (d *DBImpl) DeleteNote(ctx context.Context, id string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return true, fmt.Errorf("bad object id: %s", id)
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	_, err = d.notesCollection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *DBImpl) UpdateNote(ctx context.Context, note *Note) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(note.ID)
	if err != nil {
		return true, fmt.Errorf("bad object id: %s", note.ID)
	}

	update := bson.D{{Key: "$set", Value: note.toBSOND()}}

	_, err = d.notesCollection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
