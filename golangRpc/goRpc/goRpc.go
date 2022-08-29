package goRpc

import (
	"fmt"
	"log"
	"net/rpc"
)

/**
调用Hello的rpc方法
*/
func RpcCallHello() {

	/**
	首先是通过 rpc.Dial 拨号 RPC 服务，然后通过 client.Call 调用具体的 RPC 方法。
	在调用 client.Call 时，第一个参数是用点号连接的 RPC 服务名字和方法名字，第二和第三个参数分别我们定义 RPC 方法的两个参数。
	*/
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	/**
	调用RPC方法
	*/
	var reply string
	err = client.Call("HelloService.Hello", "i am jack", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hello reply: ", reply)

}
