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
	OpenTLS = true
)

//自定义认证
type customCredential struct{}

//GetRequestMetadata实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
		// "appkey": "i am fake key",
	}, nil
}

//RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return OpenTLS
}

//customCredential其实是实现了grpc/credential包内的PerRPCCredentials接口,这个接口中的方法就是这两个

func main() {
	var err error
	var opts []grpc.DialOption //获取grpc连接的时候其实可以传切片
	if OpenTLS {               //开启了TLS认证
		//TSL连接
		cred, err := credentials.NewClientTLSFromFile("../keys/server.pem", "www.eline.com") //第二个参数是server name，就是生成证书时写的
		if err != nil {
			grpclog.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	//customCredential其实是实现了grpc/credential包内的PerRPCCredentials接口
	//WithPerRPCCredentials的参数是传PerRPCCredentials
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	conn, err := grpc.Dial(Address, opts...) //会返回ClientConn， Dial(target string, opts ...DialOption)

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

//TODO  grpc还有拦截器，网关等，有空在学
