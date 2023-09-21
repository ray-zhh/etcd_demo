package main

import (
	"context"
	"etcd_demo/common/pb"
	"etcd_demo/etcd"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	addr = flag.String("addr", "127.0.0.1:50051", "")
)

func main() {
	flag.Parse()
	etcd.CusLoadService("echo-service")
	addr := etcd.CusServiceDiscover("echo-service")
	fmt.Println(addr)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)
	CallUnaryEcho(c)
}

func CallUnaryEcho(c pb.EchoClient) {
	ctx := context.Background()
	in := pb.EchoRequest{
		Message: "hello server",
	}

	res, err := c.UnaryEcho(ctx, &in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client recv:", res.Message)
}
