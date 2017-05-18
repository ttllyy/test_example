package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "test_example/hmsg"
)

const (
	port = ":8082"
)

type server struct {}

func (s *server) HandlerMsg(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	res := &pb.Response{
		Name:		in.Name,
		Age:		in.Age,
		Score:		in.Score,
	}
	res.Age++
	res.Score++
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHmsgServer(s, &server{})
	s.Serve(lis)
}
