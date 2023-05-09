package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "getnirm/goserver/generated"

	"google.golang.org/grpc"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// response hello world
	w.WriteHeader(http.StatusOK)
	currentTime := time.Now().Format(time.RFC850)
	w.Write([]byte("hello world " + currentTime))
}

const (
	listenAddress = "0.0.0.0:8090"
)

type timeService struct {
}

func (t *timeService) GetCurrentTime(ctx context.Context, req *pb.GetCurrentTimeRequest) (*pb.GetCurrentTimeResponse, error) {
	log.Println("Got time request")
	return &pb.GetCurrentTimeResponse{CurrentTime: time.Now().String()}, nil
}

func main() {
	// const port = 8080
	fmt.Printf("Starting server on port %s", listenAddress)
	// http.HandleFunc("/", rootHandler)
	// // convert port to string
	// // start the server on port 8080
	// serverPort := ":" + strconv.Itoa(port)
	// http.ListenAndServe(serverPort, nil)

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
