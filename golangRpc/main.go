package main

import "micro_api/golangRpc/golangRpcProto"

func main() {

	// 调用Hello的rpc方法
	//golangRpc.RpcCallHello()

	// 调用Hello02的rpc方法
	//golangRpc.RpcCallHello02()

	// 调用RpcCallHello03
	//golangRpc.RpcCallHello03()

	// 使用jsonrpc
	//golangRpc.RpcCallHello04()

	//golangRpc.RpcCallHello06()

	//==================rpc+proto客户端调用============================
	//golangRpcProto.RpcCallHello07()

	//grpc的双向流调用
	//golangRpcProto.RpcCallHello08()

	// grpc的单向流调用
	golangRpcProto.RpcCallHello09()

}
