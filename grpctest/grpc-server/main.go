package main

import (
	"fmt"
	"net"

	pb "github.com/wbing441282413/goRPCTest/grpctest/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包
	"google.golang.org/grpc/grpclog"
)

/*
	gRPC的各服务间使用高效的protobuf协议进行RPC调用，
	gRPC默认使用protocol buffers，这是google开源的一套成熟的结构数据序列化机制（当然也可以使用其他数据格式如JSON）
	可以用protocol buffers的proto文件来创建gRPC服务，用message类型来定义方法参数和返回类型

*/

// 生成了pb.go文件后，服务端就可以实现pb中的接口了
type ProgrammerServiceServer struct{}

func (p *ProgrammerServiceServer) GetProgrammerInfo(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	name := req.Name
	if name == "wb" {
		resp = &pb.Response{
			Uid:      6,
			UserName: name,
			Job:      "go",
			GoodAt:   []string{"Go", "Java", "PHP", "Python"},
		}

	}
	err = nil
	return
}

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8078"
)

//对grpc服务加上TLS认证机制，认证不是客户登录那种认证，而是是否有权利进行rpc的认证，
//需要准备好证书，这里放在keys包下，需要引入"google.golang.org/grpc/credentials"认证包

func main() {
	//实现pb中的接口
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n with TLS", Address)

	//TSL认证
	cred, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	//实例化grpc server服务器,并开启TLS认证
	// s := grpc.NewServer()
	s := grpc.NewServer(grpc.Creds(cred))

	// 将 ProgrammerService 注册到 gRPC
	// 注意第二个参数 ProgrammerServiceServer 是接口类型的变量
	// 需要取地址传参
	pb.RegisterProgrammerServiceServer(s, &ProgrammerServiceServer{}) // 注册ProgrammerServiceServer,因为接口绑定指针，所以要传指针

	s.Serve(listen)
}
