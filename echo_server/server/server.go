package server

import (
	"context"
	"etcd_demo/common/pb"
	"fmt"
)

type EchoService struct {
	pb.UnimplementedEchoServer
}

func (*EchoService) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Printf("server recv :%s\n", req.Message)
	return &pb.EchoResponse{
		Message: "hello client",
	}, nil
}
