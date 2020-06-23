package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
	"todo/config"
	"todo/list"
	"todo/model"
	"todo/pool"
	"todo/proto"
	"todo/task"
	"todo/user"
)

func main() {
	config.InitConfig()
	if config.Settings.GetBool("AUTO_MIGRATE") == true {
		model.AutoMigrate()
	}

	pool.InitDispatcher()
	pool.DispatcherInstance.Submit(pool.Job{
		ID:          1,
		Name:        "Check Deadline",
		Run:         task.NewTaskService().CheckDeadline,
		Interval:    time.Duration(60) * time.Second,
		IsRecurring: true,
	})

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
