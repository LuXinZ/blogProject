package main

import (
	"context"
	"github.com/LuXinZ/blogProject/someTest/grpcTest/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	proto.UnimplementedHelloServer
}

func (s *Server) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.Response, error) {
	return &proto.Response{
		Reply: request.Name + "hello",
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterHelloServer(g, &Server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		return
	}
}
