package golangRpcProto

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_api/micro_proto/hello"
)

func RpcCallHello09() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := hello.NewHelloServiceClient(conn)
	// 创建发送结构体
	res, err := grpcClient.ChannelOneWay(context.Background())
	if err != nil {
		log.Fatalf("Call SayHello err: %v", err)
	}
	for i := 0; i < 5; i++ {
		//通过 Send方法发送流信息
		err = res.Send(&hello.StringDto{Value: "客户端流式"})
		if err != nil {
			return
		}
	}

	// 打印返回值
	recv, err := res.CloseAndRecv()
	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
	log.Println("打印返回值: ", string(recv.Value))
}
