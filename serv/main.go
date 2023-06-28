package main

import (
	"log"
	"net"

	// Импотртируем код протобуфера
	pb "GoS/transport"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port        = ":50051"
	temperature = 25.7
	id          = 2
	name        = "Москва"
)

type server struct {
	pb.UnimplementedSendServiceServer
}

//   - мы создали только один метод для нашего сервиса,
//
// который является методом create, который принимает контекст и запрос
// потом они обрабатываются сервером gRPC.
func (s *server) SendMSG(ctx context.Context, in *pb.ClientMSG) (*pb.ServMSG, error) {

	log.Printf("Received: %s", in.GetName())
	return &pb.ServMSG{Id: id, Temperature: temperature, Name: name}, nil

}

func main() {

	// Стартуем наш gRPC сервер для прослушивания tcp
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterSendServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
