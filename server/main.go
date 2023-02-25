package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct{}

// 业务场景
// 该函数需要被客户端调用
func (s *HelloService) Hello(name string, resq *string) error {
	*resq = fmt.Sprintf("hello %s", name)
	return nil
}

func main() {
	// 服务注册给RPC框架
	err := rpc.RegisterName("HelloService", new(HelloService))
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
		
		go rpc.ServeConn(conn)
	}
}
