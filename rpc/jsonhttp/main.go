package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const (
	Name = "HelloService"
)

// http rpc

type Service interface {
	Hello(name string, resp *string) error
}

type HelloService struct{}

var _ Service = (*HelloService)(nil)

func (s *HelloService) Hello(name string, resq *string) error {
	*resq = fmt.Sprintf("hello %s", name)
	return nil
}

type RPCReadWriteCloser struct {
	io.Writer
	io.ReadCloser
}

func NewRPCReadWriteCloserFromHTTP(w http.ResponseWriter, r *http.Request) *RPCReadWriteCloser {
	return &RPCReadWriteCloser{
		w,
		r.Body,
	}
}

func httprpcService(w http.ResponseWriter, r *http.Request) {
	conn := NewRPCReadWriteCloserFromHTTP(w, r)
	rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	http.HandleFunc("/jsonrpc", httprpcService)
	http.ListenAndServe(":1234", nil)
}

/*

请求方法json
{
    "method" : "HelloService.Hello",
    "params": ["dengliyao2"],
    "id":1
}

返回
{"id":1,"result":"hello dengliyao2","error":null}

*/
