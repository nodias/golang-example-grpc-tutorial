package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"

	pb "github.com/nodias/golang-example-grpc-tutorial/grpc/v1.0"
)

const (
	address = "127.0.0.1:50051"
	defaultName ="world"
)

func main(){
	//target, dialOption을 받아 clientConn 생성
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//clientConn 객체를 protobuff의 service의 구현이 담긴 greeterClient 구조체에 담아 리턴한다.
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) >1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//call
	r, err := c.SayHello(ctx, &pb.HelloRequest{ Name:name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Printf("Greet", r.Message)
}
