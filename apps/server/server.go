package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "getnirm/goserver/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	listenAddress = "127.0.0.1:8090"
)

type timeService struct {
}

func (t *timeService) GetCurrentTime(ctx context.Context, req *pb.GetCurrentTimeRequest) (*pb.GetCurrentTimeResponse, error) {
	log.Println("Got time request")
	grpc.SendHeader(ctx, metadata.New(map[string]string{"Access-Control-Allow-Origin": "*"}))
	return &pb.GetCurrentTimeResponse{CurrentTime: time.Now().String()}, nil
}

func main() {
	fmt.Printf("Starting server on port %s", listenAddress)

	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTimeServiceServer(s, &timeService{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
