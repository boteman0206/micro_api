package golangRpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_product/micro_proto/pc"
)

func RpcCallHello07() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pc.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

}
