package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// JSON编码时参数结构体可以和服务端不一样
// 但是结构体里的字段必须一样
type A struct {
	x, y string
}

func main() {
	a := A{x: "hai", y: "oo"}

	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8086")
	if err != nil {
		log.Fatal("建立连接失败: ", err)
	}

	var result string
	err1 := client.Call("Appender.AppenderString", a, &result)
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Println("结果：", result)

}
