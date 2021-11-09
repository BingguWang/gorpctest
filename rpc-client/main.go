package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type ClientParam struct {
	X, Y int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8808") //建立rpc连接，返回client对象,通过client来进行rpc调用
	if err != nil {
		log.Fatal("dailhttp err: ", err)
	}
	fmt.Printf("%s\n", client)
	p1 := 1
	p2 := 2
	params := ClientParam{X: p1, Y: p2}
	var result int
	err = client.Call("Algorithm.Sum", &params, &result) // Call(serviceMethod string, args interface{}, reply interface{}) //服务方法名，
	if err != nil {
		log.Fatal("远程调用Sum方法错误： ", err)
	}
	fmt.Println("结果是： ", result)
}
