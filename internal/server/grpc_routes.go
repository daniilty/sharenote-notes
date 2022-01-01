package server

import (
	"context"

	schema "github.com/daniilty/sharenote-grpc-schema"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GRPC) AddNote(ctx context.Context, req *schema.AddNoteRequest) (*schema.AddNoteResponse, error) {
	note := convertPBNoteToCore(req.GetNote())

	err := g.service.AddNote(ctx, note)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.AddNoteResponse{}, nil
}

func (g *GRPC) GetNote(ctx context.Context, req *schema.GetNoteRequest) (*schema.GetNoteResponse, error) {
	note, ok, err := g.service.GetNote(ctx, req.GetId())
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.GetNoteResponse{
		Note: convertCoreNoteToPB(note),
	}, nil
}

func (g *GRPC) GetNotes(ctx context.Context, req *schema.GetNotesRequest) (*schema.GetNotesResponse, error) {
	notes, err := g.service.GetNotes(ctx, req.GetIds())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.GetNotesResponse{
		Notes: convertCoreNotesToPB(notes),
	}, nil
}

func (g *GRPC) DeleteNote(ctx context.Context, req *schema.DeleteNoteRequest) (*schema.DeleteNoteResponse, error) {
	ok, err := g.service.DeleteNote(ctx, req.GetId())
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.DeleteNoteResponse{}, nil
}

func (g *GRPC) DeleteNotes(ctx context.Context, req *schema.DeleteNotesRequest) (*schema.DeleteNotesResponse, error) {
	err := g.service.DeleteNotes(ctx, req.GetIds())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.DeleteNotesResponse{}, nil
}

func (g *GRPC) UpdateNote(ctx context.Context, req *schema.UpdateNoteRequest) (*schema.UpdateNoteResponse, error) {
	ok, err := g.service.UpdateNote(ctx, convertPBNoteToCore(req.GetNote()))
	if err != nil {
		if ok {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &schema.UpdateNoteResponse{}, nil
}
