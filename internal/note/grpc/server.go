package grpc

import (
	"context"
	"time"

	pb "github.com/deimossy/noteverse/api/proto"
)

type NoteServiceServer struct {
	pb.UnimplementedNoteServiceServer
}

func (s *NoteServiceServer) CreateNote(
	ctx context.Context,
	req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	note := &pb.Note{
		Id:        "generated-unique-id",
		UserId:    req.UserId,
		Content:   req.Content,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	return &pb.CreateNoteResponse{Note: note}, nil
}
