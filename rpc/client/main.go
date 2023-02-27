package main

import (
	"fmt"
	"net/rpc"
	
	"github.com/dengliyao/grpc-demo/rpc/service"
)

var _ service.Service = (*HelloClient)(nil)

// 客户端构造函数
func NewHelloClient(network, address string) (service.Service, error) {
	// 与服务端建立连接
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloClient{
		client: client,
	}, nil
	
}

type HelloClient struct {
	client *rpc.Client
}

// Hello 对于RPC客户端，需要包装客户端的调用
func (c *HelloClient) Hello(name string, resp *string) error {
	
	return c.client.Call(service.Name+".Hello", name, resp)
}

func main() {
	
	// 初始化客户端实例
	client, err := NewHelloClient("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	
	// 发起客户端调用
	var resp string
	client.Hello("dengliyao", &resp)
	
	fmt.Println(resp)
	
}
