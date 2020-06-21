package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"todo/list"
	"todo/proto"
	"todo/task"
	"todo/user"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	l := list.Server{}
	t := task.Server{}
	u := user.Server{}

	proto.RegisterListServiceServer(grpcServer, &l)
	proto.RegisterTaskServiceServer(grpcServer, &t)
	proto.RegisterUserServicesServer(grpcServer, &u)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
