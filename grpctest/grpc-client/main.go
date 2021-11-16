package main

import (
	"fmt"
	"log"

	pb "github.com/wbing441282413/goRPCTest/grpctest/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8078"
)

func main() {
	//TSL连接
	cred, err := credentials.NewClientTLSFromFile("../keys/server.pem", "www.eline.com") //第二个参数是server name，就是生成证书时写的

	// conn, err := grpc.Dial(Address, grpc.WithInsecure()) //会返回ClientConn
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(cred)) //会返回ClientConn
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
		grpclog.Fatalln(err)

	}

	defer conn.Close()

	// 实例化 ProgrammerService。用ClientConn对象来调用服务
	client := pb.NewProgrammerServiceClient(conn)

	// 调用服务
	req := new(pb.Request)
	req.Name = "wb"
	resp, err := client.GetProgrammerInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	fmt.Printf("Recevied: %v\n", resp)
}
