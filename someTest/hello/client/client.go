package main

import (
	"log"
	"net/rpc"
)

func main() {
	dial, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		return
	}
	var reply string
	err = dial.Call("HelloService.Hello", "jack11", &reply)
	if err != nil {
		log.Fatal("faile")
	}
	println(reply)
}
