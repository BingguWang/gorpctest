package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/wbing441282413/goRPCTest/grpctest/proto"
	"google.golang.org/grpc"
)

/*
	gRPC的各服务间使用高效的protobuf协议进行RPC调用，
	gRPC默认使用protocol buffers，这是google开源的一套成熟的结构数据序列化机制（当然也可以使用其他数据格式如JSON）
	可以用protocol buffers的proto文件来创建gRPC服务，用message类型来定义方法参数和返回类型

*/

//生成了pb.go文件后，服务端就可以实现pb中的接口了
type ProgrammerServiceServer struct{}

func (p *ProgrammerServiceServer) GetProgrammerInfo(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "shirdon" {
		resp = &pb.Response{
			Uid:      6,
			Username: name,
			Job:      "CTO",
			GoodAt:   []string{"Go", "Java", "PHP", "Python"},
		}

	}
	err = nil
	return
}

func main() {
	//实现pb中的接口
	port := ":8078"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()
	// 将 ProgrammerService 注册到 gRPC
	// 注意第二个参数 ProgrammerServiceServer 是接口类型的变量
	// 需要取地址传参
	pb.RegisterProgrammerServiceServer(s, &ProgrammerServiceServer{})
	s.Serve(l)
}
