package main

import (
	pb "Go-Grpc/proto"
	"context"
)

func (s *helloServer) sayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
