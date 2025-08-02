package grpc

import (
	"context"

	pb "github.com/deimossy/noteverse/api/proto"

)

type NoteServiceServer struct {
	pb.UnimplementedNoteServiceServer
}

func (s *NoteServiceServer) CreateNote(ctx context.Context)