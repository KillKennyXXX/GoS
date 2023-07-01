package main

import (
	"log"
	"net"

	av "GoS/Avro/avroSchema"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	av.UnimplementedTestServer
}

//   - мы создали только один метод для нашего сервиса,
//
// который является методом create, который принимает контекст и запрос
// потом они обрабатываются сервером gRPC.

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
