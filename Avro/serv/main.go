package main

import (
	av "GoS/Avro/avroSchema"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
)

type server struct {
	av.UnimplementedTestServer
}

func (s *server) Hello(ctx context.Context, in map[string]interface{}) (interface{}, error) {
	name := in["name"]
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello %s!", name), nil
}

func main() {

	// Стартуем наш gRPC сервер для прослушивания tcp
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	av.RegisterTestServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
