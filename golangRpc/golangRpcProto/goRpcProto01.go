package golangRpcProto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_api/micro_proto/pc"
)

/**
gRPC 和标准库的 RPC 框架有一个区别，gRPC 生成的接口并不支持异步调用

*/

func RpcCallHello07() {

	//其中 grpc.Dial 负责和 gRPC 服务建立连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//然后 NewHelloServiceClient 函数基于已经建立的连接构造 HelloServiceClient 对象
	client := pc.NewHelloServiceClient(conn)
	// 返回的 client 其实是一个 HelloServiceClient 接口对象，通过接口定义的方法就可以调用服务端对应的 gRPC 服务提供的方法。
	reply, err := client.Hello(context.Background(), &pc.StringDto{Value: "proto-hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

}
