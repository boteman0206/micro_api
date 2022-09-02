package goRpcAdvance

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_api/micro_proto/hello"
)

/**
grpc的拦截器的使用
*/
func RpcCallHello14() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hello.StringDto{Value: "interceptor-hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

}
