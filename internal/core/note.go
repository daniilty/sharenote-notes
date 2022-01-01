package core

import (
	"context"
)

type Note struct {
	ID        string
	Data      string
	Name      string
	Timestamp int64
}

func (s *ServiceImpl) AddNote(ctx context.Context, note *Note) error {
	return s.db.AddNote(ctx, note.toDB())
}

func (s *ServiceImpl) GetNote(ctx context.Context, id string) (*Note, bool, error) {
	note, ok, err := s.db.GetNote(ctx, id)
	if err != nil {
		return nil, ok, err
	}

	return convertDBNoteToService(note), true, nil
}

func (s *ServiceImpl) GetNotes(ctx context.Context, ids []string) ([]*Note, error) {
	notes := make([]*Note, 0, len(ids))

	for i := range ids {
		note, ok, err := s.db.GetNote(ctx, ids[i])
		if err != nil {
			if ok {
				continue
			}

			return nil, err
		}

		notes = append(notes, convertDBNoteToService(note))
	}

	return notes, nil
}

func (s *ServiceImpl) DeleteNote(ctx context.Context, id string) (bool, error) {
	return s.db.DeleteNote(ctx, id)
}

func (s *ServiceImpl) DeleteNotes(ctx context.Context, ids []string) error {
	for i := range ids {
		_, err := s.db.DeleteNote(ctx, ids[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ServiceImpl) UpdateNote(ctx context.Context, note *Note) (bool, error) {
	return s.db.UpdateNote(ctx, note.toDB())
}
