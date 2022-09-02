package main

import (
	"micro_api/golangRpc/goRpcAdvance"
	//"micro_api/golangRpc/golangRpcProto"
)

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
	//golangRpcProto.RpcCallHello09()

	// 发布-订阅模型使用

	//go func() {
	//	golangRpcProto.SubscribeMsg()
	//}()
	//
	//time.Sleep(10 * time.Second)
	//golangRpcProto.PublishMsg()
	//
	//for true {
	//
	//}

	//==============goRpc的Tls调用===========================
	//自签名证书调用
	//goRpcAdvance.RpcCallHello11()

	//根证书签名测试
	//goRpcAdvance.RpcCallHello12()

	// token认证的rpc
	//goRpcAdvance.RpcCallHello13()

	//拦截器的使用
	goRpcAdvance.RpcCallHello14()

}
