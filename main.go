package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"todo/list"
	"todo/model"
	"todo/proto"
	"todo/task"
	"todo/user"
)

func main() {
	AutoMigrate()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	l := list.Server{}
	t := task.Server{}
	proto.RegisterListServiceServer(grpcServer, &l)
	proto.RegisterTaskServiceServer(grpcServer, &t)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func AutoMigrate() {
	model.Client().AutoMigrate(&user.User{}, &list.List{}, &task.Task{})
}
