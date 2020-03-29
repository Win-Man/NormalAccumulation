package main

import (
	"context"
	"io"
	"fmt"
	"log"
	"net"

	pb "github.com/Win-Man/sg-server/protos"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received Default request from: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server)TellMeSomething(in *pb.HelloRequest, stream pb.Greeter_TellMeSomethingServer) error{
	log.Printf("Received ServerStream request from: %v", in.GetName())
	// 获取客户端发送的请求内容
	hellostr := fmt.Sprintf("Hello,%s",in.GetName())
	something := []string{hellostr,"ServerLine1","ServerLine2","ServerLine3"}
	for i,v := range(something){
		if err := stream.Send(&pb.Something{LineCode:int64(i),Line:v});err != nil{
			return err
		}
	}
	return nil
}

func (s *server)TellYouSomething(stream pb.Greeter_TellYouSomethingServer) error{
	log.Printf("Recevied ClientStream request")
	messgeCount := 0
	length := 0
	for {
		something,err := stream.Recv()
		// 接收到客户端所有请求，返回响应结果
		if err == io.EOF{
			return stream.SendAndClose(&pb.HelloReply{Message:fmt.Sprintf("It's over.\nReceived %v times. Length:%v",messgeCount,length)})
		}
		if err != nil{
			return err
		}
		messgeCount ++
		length += len(something.GetLine())
	}
	return nil
}

func (s *server)TalkWithMe(stream pb.Greeter_TalkWithMeServer) error{
	log.Printf("Recevied Stream request")
	messageCount := 0
	for {
		something ,err := stream.Recv()
		if err == io.EOF{
			return nil
		}
		messageCount ++
		length := len(something.GetLine())
		line := fmt.Sprintf("Got %s,Length:%v",something.GetLine(),length)
		if err := stream.Send(&pb.Something{LineCode: int64(messageCount),Line:line});err != nil{
			return err
		}
	}
	return nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
