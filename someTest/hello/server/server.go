package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
func main() {
	// new server
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	// register handle
	err = rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		return
	}
	// start server
	conn, err := listen.Accept()
	rpc.ServeConn(conn)
}
