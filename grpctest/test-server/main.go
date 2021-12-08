package main

import (
	"fmt"
	pb "github.com/wbing441282413/goRPCTest/grpctest/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const addr string = "127.0.0.1:22000"

func main() {
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, HelloServiceImpl{})

	s.Serve(listen)
}

type HelloServiceImpl struct{}

func (*HelloServiceImpl) GetHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println(req)
	resp := &pb.HelloResponse{Msg: "sad"}
	return resp, nil
}
