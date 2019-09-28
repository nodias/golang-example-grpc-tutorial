package client

import (
	"google.golang.org/grpc"
	"log"

	pb "github.com/nodias/golang-example-grpc-tutorial/grpc/v1"
)

const (
	address = "localhost:50051"
	defaultName ="world"
)

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

}
