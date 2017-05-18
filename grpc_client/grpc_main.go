package main

import (
	"log"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "test_example/hmsg"

)

const (
	address     = "127.0.0.1:8082"
)




func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewHmsgClient(conn)
	request := &pb.Request{
		Name: 	"qiaoyang",
		Age:	77,
		Score:	90,
	}

	r, err := c.HandlerMsg(context.Background(), request)
	if err != nil {
		log.Fatal("could not saymsg: %v", err)
	}

	fmt.Println(*r)

	fmt.Println("exit")
}

