package main

import (
	"etcd_demo/common/pb"
	"etcd_demo/echo_server/server"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server.EchoService{})

	//err = etcd.CusServiceRegister("echo-service", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
