package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//服务端，实现用JSON编码数据的rpc服务器

type StringApp struct {
	x, y string
}

type Appender string //建立服务对象，

func (appender *Appender) AppenderString(args StringApp, reply *string) error {
	*reply = args.x + args.y
	fmt.Println(reply)
	return nil
}

//使用Go提供的net/rpc/jsonrpc标准包
func init() {
	fmt.Println("JSON编码RPC，不是gob编码，其他的和RPC概念一模一样，")
}
func main() {
	//实例化服务对象
	appender := new(Appender)
	//注册
	rpc.Register(appender)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8086")
	if err != nil {
		fmt.Println("ResolveTCPAddr err=", err)
	}
	listen, err := net.ListenTCP("tcp", tcpAddr) //目前jsonrpc库是基于tcp协议实现的，暂不支持http传输方式
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err1 := listen.Accept()
		if err1 != nil {
			log.Fatal(err1)
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
