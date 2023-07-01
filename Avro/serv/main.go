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

type testServer struct {
	av.UnimplementedTestServer
}

func (s *testServer) Hello(ctx context.Context, in map[string]interface{}) (interface{}, error) {
	name := in["name"]
	if name == "" {
		name = "World"
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
	av.RegisterTestServer(s, &testServer{})
	go s.Serve(lis)
	// defer s.GracefulStop()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
