package main

import (
	"flag"
	"log"
	"time"

	// Импотртируем код протобуфера
	av "GoS/Avro/avroSchema"

	// "encoding/json"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	data = map[string]interface{}{"name": ""}
)

// func (c *testClient) Hello(ctx context.Context, in map[string]interface{}, opts ...grpc.CallOption) (interface{}, error) {
// 	buf, err := _Test_Hello_requestCodec.BinaryFromNative(nil, in)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = c.cc.Invoke(ctx, "/test/hello", buf, &buf, opts...); err != nil {
// 		return nil, err
// 	}
// 	out, _, err := _Test_Hello_responseCodec.NativeFromBinary(buf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return out, nil
// }

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := av.NewTestClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Hello(ctx)
	// r, err := c.Hello(ctx, *&data, grpc.CallContentSubtype("avro"))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// j, _ := json.Marshal(r)
	// // log.Printf("вы в городе %s, температура  %f*, номер сообщения %d ", r.GetName(), r.GetTemperature(), r.GetId())

	log.Printf("%s", r)
}
