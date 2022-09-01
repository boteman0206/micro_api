package goRpcAdvance

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_api/micro_proto/hello"
)

/**
使用token认证的grpc的服务端
type PerRPCCredentials interface
要实现对每个 gRPC 方法进行认证，需要实现 grpc.PerRPCCredentials 接口：
*/

type Authentication struct {
	User     string
	Password string
}

//我们可以创建一个 Authentication 类型，用于实现用户名和密码的认证：
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

//要实现对每个 gRPC 方法进行认证，需要实现 grpc.PerRPCCredentials 接口：
/**
在 GetRequestMetadata 方法中，我们返回地认证信息包装 user 和 password 两个信息。为了演示代码简单，RequireTransportSecurity 方法表示不要求底层使用安全连接。
然后在每次请求 gRPC 服务时就可以将 Token 信息作为参数选项传人：
*/

func RpcCallHello13() {
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}

	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := hello.NewAuthServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hello.StringDto{Value: "token-hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

}
