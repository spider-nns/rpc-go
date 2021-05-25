package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	pb "rpc-go/proto"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}
func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "spider"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}
func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
}
