package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}
func main() {
	// new server
	http.HandleFunc("/jsporpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer:     writer,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	// register handle
	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		return
	}
	http.ListenAndServe(":1234", nil)
}
