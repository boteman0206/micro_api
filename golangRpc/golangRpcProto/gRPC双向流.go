package golangRpcProto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"micro_api/micro_proto/hello"
	"time"
)

func RpcCallHello08() {

	//其中 grpc.Dial 负责和 gRPC 服务建立连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)

	//客户端需要先调用 Channel 方法获取返回的流对象：
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//在客户端我们将发送和接收操作放到两个独立的 Goroutine。
	// 首先是向服务端发送数据：
	go func() {
		for {
			if err := stream.Send(&hello.StringDto{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	// 然后在循环中接收服务端返回的数据：
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println("客户端获取的值： ", reply.GetValue())
	}

	time.Sleep(30 * time.Minute)
}
