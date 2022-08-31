package golangRpcProto

import (
	"google.golang.org/grpc"
	"log"
)

func RpcCallHello08() {

	//其中 grpc.Dial 负责和 gRPC 服务建立连接
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

}
