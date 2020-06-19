package task

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

	It("Get Single Task By UUID", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()

		client := proto.NewTaskServiceClient(conn)

		task := &model.Task{}
		err = model.Client().Model(task).First(&task).Error
		resp, err := client.GetTaskByID(ctx, &proto.TaskRequest{TaskID: task.UUID})
		Expect(err).To(BeNil())
		Expect(resp.ID).To(BeEquivalentTo(task.UUID))
	})

	It("Get Tasks For User", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()
		client := proto.NewTaskServiceClient(conn)

		user := &u.User{}
		err = model.Client().Model(user).First(&user).Error

		q := model.Client().Model(&model.Task{}).Where("user_id = ?", user.ID)

		tasks := []*model.Task{}

		err = q.Find(&tasks).Error

		resp, err := client.GetTasksForUser(ctx, &proto.UserTasksRequest{UserID: user.UUID})

		Expect(err).To(BeNil())
		Expect(len(resp.Tasks)).To(Not(BeEquivalentTo(0)))
		Expect(len(resp.Tasks)).To(BeEquivalentTo(len(tasks)))
		for _, v := range resp.Tasks {
			Expect(v.UserID).To(BeEquivalentTo(user.UUID))
		}
	})

	It("Get Tasks For List", func() {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(BufferDialer), grpc.WithInsecure())
		defer conn.Close()
		client := proto.NewTaskServiceClient(conn)

		l := &model.List{}
		err = model.Client().Model(l).First(&l).Error

		q := model.Client().Model(&model.Task{}).Where("list_id = ?", l.ID)

		tasks := []*model.Task{}

		err = q.Find(&tasks).Error

		resp, err := client.GetTasksForList(ctx, &proto.ListTasksRequest{ListID: l.UUID})

		Expect(err).To(BeNil())
		Expect(len(resp.Tasks)).To(Not(BeEquivalentTo(0)))
		Expect(len(resp.Tasks)).To(BeEquivalentTo(len(tasks)))
		for _, v := range resp.Tasks {
			Expect(v.ListID).To(BeEquivalentTo(l.UUID))
		}
	})
})
