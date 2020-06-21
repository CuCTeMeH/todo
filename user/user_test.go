package user

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"todo/config"
	"todo/model"
	"todo/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func BufferDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func PrepareServer() {
	config.InitConfig()
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

	It("Create User", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()
		client := proto.NewUserServicesClient(conn)

		user := &model.User{}
		err = model.Client().Model(user).First(&user).Error

		l := &model.List{}
		err = model.Client().Model(l).Where("user_id = ?", user.ID).First(&l).Error

		resp, err := client.NewUser(ctx, &proto.NewUserRequest{Username: "test_username", Email: "test_email@email.com", FirstName: "Test First Name", LastName: "Test Last Name"})

		Expect(err).To(BeNil())
		Expect(resp.Username).To(BeEquivalentTo("test_username"))
		Expect(resp.Email).To(BeEquivalentTo("test_email@email.com"))
		Expect(resp.FirstName).To(BeEquivalentTo("Test First Name"))
		Expect(resp.LastName).To(BeEquivalentTo("Test Last Name"))

		//delete from db after test
		err = model.Client().Model(&model.User{}).Where("uuid = ?", resp.ID).Unscoped().Delete(&model.User{}).Error
		Expect(err).To(BeNil())
	})

	It("Edit User", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()
		client := proto.NewUserServicesClient(conn)

		user := &model.User{}
		err = model.Client().Model(user).First(&user).Error

		resp, err := client.EditUser(ctx, &proto.EditUserRequest{UserID: user.UUID, User: &proto.NewUserRequest{Username: "edit_username", Email: "edit_test_email@gmail.com", FirstName: "Edit Test First Name", LastName: "Edit Test Last Name"}})

		Expect(err).To(BeNil())
		Expect(resp.ID).To(BeEquivalentTo(user.UUID))

		//delete from db after test
		resp, err = client.EditUser(ctx, &proto.EditUserRequest{UserID: user.UUID, User: &proto.NewUserRequest{Username: user.Username, Email: user.Email, FirstName: user.FirstName, LastName: user.LastName}})
		Expect(err).To(BeNil())
	})

})
