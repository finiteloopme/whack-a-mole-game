package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/finiteloopme/whack-a-mole-game/mole"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMoleServer
}

// rpc GetMoles(stream MoleFromClient) returns (stream MoleFromServer);
func (s *server) GetMoles(stream pb.Mole_GetMolesServer) error {
	log.Printf("Received: ")
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoleServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

