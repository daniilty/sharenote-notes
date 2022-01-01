package server

import (
	schema "github.com/daniilty/sharenote-grpc-schema"
	"github.com/daniilty/sharenote-notes/internal/core"
)

type GRPC struct {
	schema.UnimplementedNotesServer

	service core.Service
}

func NewGRPC(noteService core.Service) *GRPC {
	return &GRPC{
		service: noteService,
	}
}
