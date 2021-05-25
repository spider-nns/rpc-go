package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "rpc-go/proto"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type GreeterServer struct {
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("receive msg:%s", r.Name)
	return &pb.HelloResponse{Message: "hello,world"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
