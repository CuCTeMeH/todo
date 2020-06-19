package task

import (
	"context"
	. "github.com/onsi/ginkgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"todo/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func BufferDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func PrepareServer() {
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	s := Server{}
	proto.RegisterTaskServiceServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

var _ = Describe("Task methods", func() {
	BeforeSuite(func() {
		PrepareServer()
	})

})
