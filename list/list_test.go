package list

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
	u "todo/user"
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
	proto.RegisterListServiceServer(grpcServer, &s)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

var _ = Describe("Lists methods", func() {
	BeforeSuite(func() {
		PrepareServer()
	})

	It("Get List By UUID", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()

		client := proto.NewListServiceClient(conn)

		list := &model.List{}
		err = model.Client().Model(list).First(&list).Error

		resp, err := client.GetListByID(ctx, &proto.ListRequest{ListID: list.UUID})
		Expect(err).To(BeNil())
		Expect(resp.ID).To(BeEquivalentTo(list.UUID))
	})

	It("Call Lists By User UUID", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()
		client := proto.NewListServiceClient(conn)

		user := &u.User{}
		err = model.Client().Model(user).First(&user).Error

		q := model.Client().Model(&model.List{}).Where("user_id = ?", user.ID)

		lists := []*model.List{}

		err = q.Find(&lists).Error

		resp, err := client.GetListsForUser(ctx, &proto.UserListsRequest{UserID: user.UUID})

		Expect(err).To(BeNil())
		Expect(len(resp.Lists)).To(Not(BeEquivalentTo(0)))
		Expect(len(resp.Lists)).To(BeEquivalentTo(len(lists)))
		for _, v := range resp.Lists {
			Expect(v.UserID).To(BeEquivalentTo(user.UUID))
		}
	})
})
