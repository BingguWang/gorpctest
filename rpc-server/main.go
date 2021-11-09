package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

//RPC，远程调用，就是不同的服务之间，或不同的计算机之间可以相互调用方法，相互调用资源，常用于微服务中

//服务端，实现一个RPC服务器端
//标准库中是有rpc的包的，先来看看使用这个包如何来做的RPC服务器
//这个的rpc包是有不足的：服务器端可以注册多个不同类型对象，但是不能注册多个同类型的对象

type Args struct {
	X, Y int
}

type Algorithm int //定义一个服务对象

//定义服务对象的方法
//第一个参数是调用者提供的参数，第二个参数是返回结果，第二个参数必须是指针！！
func (t *Algorithm) Sum(args *Args, reply *int) error {
	//注意函数签名必须满足这里的格式才能实现RPC（这就是不足之处）
	*reply = args.X + args.Y
	fmt.Println(reply)
	return nil
}

func main() {
	//1 实例化服务对象
	algorithm := new(Algorithm)
	//2 注册服务对象
	rpc.Register(algorithm)
	rpc.HandleHTTP()                         //通过http暴露出此服务
	err := http.ListenAndServe(":8808", nil) //监听并提供服务
	if err != nil {
		log.Fatal(err)
	}

}
