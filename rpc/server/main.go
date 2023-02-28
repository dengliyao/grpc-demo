package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	
	"github.com/dengliyao/grpc-demo/rpc/service"
)

// 使用变量声明来约束 HelloService 使用规范
// 声明了一个空指针, 强制把这个指针转换成一个*HelloService
var _ service.Service = (*HelloService)(nil)

type HelloService struct{}

// 业务场景
// 该函数需要被客户端调用
func (s *HelloService) Hello(name string, resq *string) error {
	*resq = fmt.Sprintf("hello %s", name)
	return nil
}

// tcprpc
func main() {
	// 服务注册给RPC框架
	err := rpc.RegisterName(service.Name, new(HelloService))
	if err != nil {
		panic(err)
	}
	
	// 监听socket
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	
	// 处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		// 实现 服务端json
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
