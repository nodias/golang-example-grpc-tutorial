package main

import (
	pb "github.com/nodias/golang-example-grpc-tutorial/grpc/v1.0"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {
	//net.Listener interface 객체 생성
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Create *grpc.Server
	s:= grpc.NewServer()
	//grpc.Server에 ptotobuff의 service를 구현한 server를 등록한다.
	pb.RegisterGreeterServer(s, &server{})
	//net.Listener객체를 이용해 서비스 시작
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %v", err)
	}
}
