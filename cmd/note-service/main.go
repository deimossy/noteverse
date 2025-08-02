package main

import (
	"log"
	"net"
	
	"google.golang.org/grpc"

	pb "github.com/deimossy/noteverse/api/proto"
	notegrpc  "github.com/deimossy/noteverse/internal/note/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	noteService := &notegrpc.NoteServiceServer{}

	pb.RegisterNoteServiceServer(s, noteService)

	log.Println("gRPC server listening at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
