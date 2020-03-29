package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/Win-Man/sg-server/protos"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	clientFunc := "Default"
	if len(os.Args) == 3 {
		name = os.Args[1]
		clientFunc = os.Args[2]
	}else{
		log.Fatal("Please input 2 arguments")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	switch clientFunc {
	case "Default":
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("Default error: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	case "ServerStream":
		stream,err := c.TellMeSomething(ctx,&pb.HelloRequest{Name:name})
		if err != nil{
			log.Fatalf("TellMeSomething error: %v ",err)
		}
		for {
			something,err := stream.Recv()
			// 服务端信息发送完成，退出
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Fatalf("TellMeSomething stream error:%v",err)
			}
			log.Printf("Recevie from server:{LineCode:%v Line:%s}\n",something.GetLineCode(),something.GetLine())
		}
	case "ClientStream":
		stream,err := c.TellYouSomething(ctx)
		if err != nil{
			log.Fatalf("TellYouSomething err :%v",err)
		}
		clientStr := []string{"ClientLine1","ClientLine2"}
		for i,v := range(clientStr){
			if err := stream.Send(&pb.Something{LineCode:int64(i),Line:v});err != nil{
				log.Fatalf("TellYouSomething stream error:%v",err)
			}
		}
		rep,err := stream.CloseAndRecv()
		if err != nil{
			log.Fatalf("CloseAndRecd error:%v",err)
		}
		log.Printf("Recive from server:%s",rep.GetMessage())
	case "Stream":
		stream,err := c.TalkWithMe(ctx)
		if err != nil{
			log.Fatalf("TalkWithMe err:%v",err)
		}
		waitc := make(chan struct{})
		go func(){
			for {
				something,err := stream.Recv()
				// 服务端信息发送完成，退出
				if err == io.EOF{
					break
				}
				if err != nil{
					log.Fatalf("TalkWithMe stream error:%v",err)
				}
				log.Printf("Got %v:%s\n",something.GetLineCode(),something.GetLine())
			}
		}() 
		clientStr := []string{"one","two","three"}
		for i,v := range(clientStr){
			if err := stream.Send(&pb.Something{LineCode:int64(i),Line:v});err != nil{
				log.Fatalf("TalkWithMe Send error:%v",err)
			}
		}
		stream.CloseSend()
		<- waitc
	default:
		log.Fatal("Please input second args in Default/ServerStream/ClientStream/Stream")
	}
	
}
