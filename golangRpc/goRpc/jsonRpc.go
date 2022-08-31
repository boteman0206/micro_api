package goRpc

import (
	"encoding/json"
	"fmt"
	"log"
	"micro_api/micro_proto"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
实现 json 版本的客户端：
*/

//{"method":"HelloService.Hello","params":["hello"],"id":0}

type ResJson struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	Id     int      `json:"id"`
}

func RpcCallHello04() {

	var data = `{"method":"HelloService.Hello","params":["hello"],"id":0}`

	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(micro_proto.HelloServiceName+".Hello", data, &reply)
	if err != nil {
		log.Fatal(err)
	}

	var res ResJson
	err = json.Unmarshal([]byte(reply), &res)
	if err != nil {
		fmt.Println("json 转换错误： ", err.Error())
	}
	fmt.Println(reply, " json: ", res)
}
