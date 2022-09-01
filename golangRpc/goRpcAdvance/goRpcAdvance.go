package goRpcAdvance

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"micro_api/micro_proto/hello"
)

/**
在客户端基于服务器的证书和服务器名字就可以对服务器进行验证：
*/
func RpcCallHello11() {

	serverCrt := `./golangRpc/goRpcAdvance/server.crt`

	creds, err := credentials.NewClientTLSFromFile(serverCrt, "server.grpc.io")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//然后 NewHelloServiceClient 函数基于已经建立的连接构造 HelloServiceClient 对象
	client := hello.NewHelloServiceClient(conn)
	// 返回的 client 其实是一个 HelloServiceClient 接口对象，通过接口定义的方法就可以调用服务端对应的 gRPC 服务提供的方法。
	reply, err := client.Hello(context.Background(), &hello.StringDto{Value: "tls-hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

}
