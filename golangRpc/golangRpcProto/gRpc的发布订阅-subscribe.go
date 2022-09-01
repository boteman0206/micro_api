package golangRpcProto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"micro_api/micro_proto/hello"
)

/**

grpc发布订阅模式的 客户端订阅

*/
func SubscribeMsg() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewPubsubServiceClient(conn)

	// 订阅golang开头的消息
	streamGolang, err := client.Subscribe(context.Background(), &hello.StringDto{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			reply, err := streamGolang.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}

			fmt.Println("订阅消息golang消息： ", reply.GetValue())
		}
	}()

	// 订阅docker开头的消息
	streamDocker, err := client.Subscribe(context.Background(), &hello.StringDto{Value: "docker:"})
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			reply, err := streamDocker.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}

			fmt.Println("订阅消息docker消息： ", reply.GetValue())
		}
	}()

	for true {

	}
}
