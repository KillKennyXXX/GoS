package GoS

import (
	"log"
	"net"

	// Импотртируем код протобуфера
	pb "GoS/transport"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// IRepository - интерфейс хранилища
type IRepository interface {
	Create(*pb.ServMSG) (*pb.ServMSG, error)
}

// Repository - структура для эмитации хранилища,
// после мы заменим её на настоящие хранилищем
type Repository struct {
	messages []*pb.ServMSG
}

// Create - создаём новое хранилище
func (repo *Repository) Create(smsg *pb.ServMSG) (*pb.ServMSG, error) {
	updated := append(repo.messages, smsg)
	repo.messages = updated
	return smsg, nil
}

// Служба должна реализовать все методы для удовлетворения сервиса
// которые мы определили в нашем определении протобуфа. Вы можете проверить интерфейсы
// в сгенерированном коде для точных сигнатур методов и т. д.
type service struct {
	repo IRepository
}

//   - мы создали только один метод для нашего сервиса,
//
// который является методом create, который принимает контекст и запрос
// потом они обрабатываются сервером gRPC.
func (s *service) SendMSG(ctx context.Context, req *pb.ServMSG) (*pb.ClientMSG, error) {

	// Сохраним нашу партию в хранидище
	msg, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Возвращаем сообщение `Response`,
	// которое мы создали в нашем определнии пробуфа
	return &pb.ClientMSG{Created: true, Smsg: msg}, nil
}

func main() {

	repo := &Repository{}

	// Стартуем наш gRPC сервер для прослушивания tcp
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Зарегистрируйте нашу службу на сервере gRPC, это свяжет нашу
	// реализацию с кодом автогенерированного интерфейса для нашего
	// сообщения `Response`, которое мы создали в нашем протобуфе
	pb.RegisterSendServiceServer(s, &service{repo})

	// Регистрация службы ответов на сервере gRPC.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
