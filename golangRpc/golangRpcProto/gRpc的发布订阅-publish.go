package golangRpcProto

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"micro_api/micro_proto/hello"
)

/**
grpc发布订阅模式的 客户端发布
*/
func PublishMsg() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewPubsubServiceClient(conn)

	// 客户端发布golang：开头的消息
	_, err = client.Publish(context.Background(), &hello.StringDto{Value: "golang: hello Go"})
	if err != nil {
		log.Fatal(err)
	}

	// 客户端发布docker: 开头的消息
	_, err = client.Publish(context.Background(), &hello.StringDto{Value: "docker: hello Docker"})
	if err != nil {
		log.Fatal(err)
	}
}
