package user

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"todo/model"
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
	proto.RegisterUserServicesServer(grpcServer, &s)

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

	It("Get User By UUID", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()

		client := proto.NewUserServicesClient(conn)

		user := &model.User{}
		err = model.Client().Model(user).First(&user).Error

		resp, err := client.GetUserByID(ctx, &proto.UserRequest{UserID: user.UUID})
		Expect(err).To(BeNil())
		Expect(resp.ID).To(BeEquivalentTo(user.UUID))
	})

	It("Get User By Email", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()

		client := proto.NewUserServicesClient(conn)

		user := &model.User{}
		err = model.Client().Model(user).First(&user).Error

		resp, err := client.GetUserByEmail(ctx, &proto.UserByEmailRequest{Email: user.Email})
		Expect(err).To(BeNil())
		Expect(resp.Email).To(BeEquivalentTo(user.Email))
	})
})
