package main

import (
	"context"
	"fmt"
	"log"
	mv "movie/movie"
	"net"

	"google.golang.org/grpc"
)

type movieServiceServer struct {
	mv.UnimplementedMovieServiceServer
}

func (s *movieServiceServer) GetMovieInfo(cont context.Context, request *mv.GetMovieResultRequest) (*mv.GetMovieResultResponse, error) {
	movieResult := mv.GetMovieResultResponse{
		Title: "DDLJ"}
	fmt.Printf(movieResult.GetTitle())
	return &movieResult, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()
	mv.RegisterMovieServiceServer(s, &movieServiceServer{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
