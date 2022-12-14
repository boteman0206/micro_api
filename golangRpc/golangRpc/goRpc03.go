package golangRpc

import (
	"fmt"
	"log"
	"micro_api/micro_proto"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(micro_proto.HelloServiceName+".Hello", request, reply)
}

func RpcCallHello03() {

	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("who are you!", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RpcCallHello03 reply: ", reply)
}
