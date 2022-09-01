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

这种方式，需要提前将服务器的证书告知客户端，这样客户端在连接服务器时才能进行对服务器证书认证
在复杂的网络环境中，服务器证书的传输本身也是一个非常危险的问题。如果在中间某个环节，服务器证书被监听或替换那么对服务器的认证也将不再可靠。

为了避免证书的传递过程中被篡改，可以通过一个安全可靠的根证书分别对服务器和客户端的证书进行签名。这样客户端或服务器在收到对方的证书后可以通过根证书进行验证证书的有效性。
根证书的生成方式和自签名证书的生成方式类似：

$ openssl genrsa -out ca.key 2048
$ openssl req -new -x509 -days 3650 -subj "/C=GB/L=China/O=gobook/CN=github.com" -key ca.key -out ca.crt


*/

func RpcCallHello11() {

	// 使用的是和服务端一样的crt文件
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
