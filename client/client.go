package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "rpc-go/proto"
	"strconv"
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
func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("sayList.resp :%v", resp)
	}
	return nil
}
func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for i := 0; i < 6; i++ {
		r.Name = strconv.Itoa(i) + r.GetName()
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()
	log.Printf("sayRecord.resp: %v", resp)
	return nil
}
func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)
	request := pb.HelloRequest{Name: "sss"}
	_ = SayList(client, &request)
	recordReq := pb.HelloRequest{Name: "55"}
	_ = SayRecord(client, &recordReq)
}
