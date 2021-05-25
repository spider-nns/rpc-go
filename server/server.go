package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "rpc-go/proto"
	"strconv"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type GreeterService struct {
}

func (s *GreeterService) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("receive msg:%s", r.Name)
	return &pb.HelloResponse{Message: "hello,world"}, nil
}
func (s *GreeterService) SayList(req *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for i := 0; i < 6; i++ {
		var resp = strconv.Itoa(i) + ":√"
		log.Printf("server.sayList resp: %v", resp)
		_ = stream.Send(&pb.HelloResponse{Message: "hello.list"})
	}
	return nil
}
func (s *GreeterService) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		//ctx := context.Background()
		//log.Printf("greetServer.sayRecord.param:%v", ctx.Value("name"))
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{Message: "sayRecode:-> √"})
		}
		if err != nil {
			return err
		}
		log.Printf("resp :%v", resp)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterService{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
