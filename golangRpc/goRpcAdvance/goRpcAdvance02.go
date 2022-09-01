package goRpcAdvance

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"micro_api/micro_proto/hello"
)

/**
目前这个还无法调用
*/
func RpcCallHello12() {

	clientCrt := `./golangRpc/goRpcAdvance/client.crt`
	clientKey := `./golangRpc/goRpcAdvance/client.key`
	caCrt := `./golangRpc/goRpcAdvance/ca.crt`

	certificate, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(caCrt)
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	tlsServerName := "client.grpc.io"
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   tlsServerName, // NOTE: this is required!
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial("localhost:5000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//然后 NewHelloServiceClient 函数基于已经建立的连接构造 HelloServiceClient 对象
	client := hello.NewHelloServiceClient(conn)
	// 返回的 client 其实是一个 HelloServiceClient 接口对象，通过接口定义的方法就可以调用服务端对应的 gRPC 服务提供的方法。
	reply, err := client.Hello(context.Background(), &hello.StringDto{Value: "tls-github-hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
