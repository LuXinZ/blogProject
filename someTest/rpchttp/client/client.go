package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		return
	}
	var reply string
	//err = dial.Call("HelloService.Hello", "jack11", &reply)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "uzi111", &reply)
	if err != nil {
		log.Fatal("faile")
	}
	println(reply)
}
