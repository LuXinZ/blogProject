package main

import (
	"context"
	"github.com/LuXinZ/blogProject/someTest/grpcTest/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c := proto.NewHelloClient(conn)
	r, err := c.Hello(context.Background(), &proto.HelloRequest{
		Name: "jack",
	})
	if err != nil {
		panic(err)
	}
	println(r.Reply)
}
