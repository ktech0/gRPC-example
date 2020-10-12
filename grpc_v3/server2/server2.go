package main

import (
	"context"
	"fmt"
	"mathproto"
	"net"
	"time"

	"go.etcd.io/etcd/clientv3"

	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type Math struct{}

func (m *Math) Plus(ctx context.Context, req *mathproto.PlusRequest) (*mathproto.PlusResponse, error) {
	return &mathproto.PlusResponse{C: req.A - req.B}, nil
}

func main() {
	grpc.UseCompressor(gzip.Name)
	serveAddress := "127.0.0.1:5577"
	serveName := "MathService.2"
	register(serveName, serveAddress)
	serve(serveAddress)
}

func register(serveName string, serveAddress string) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	lease, err := etcdClient.Grant(context.TODO(), 2)
	if err != nil {
		panic(err.Error())
	}
	etcdClient.Put(ctx, serveName, serveAddress, clientv3.WithLease(lease.ID))
	//注意：这里的 context 一般为 TODO() 或 Background()。如果是其他类型的 context，则租约将根据 context 类型进行取消
	etcdClient.KeepAlive(context.Background(), lease.ID)
	fmt.Println("服务已注册")
}

func serve(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err.Error())
	}
	s := grpc.NewServer()
	mathproto.RegisterMathServer(s, &Math{})
	reflection.Register(s)
	fmt.Println("服务已启动")
	if err := s.Serve(lis); err != nil {
		panic(err.Error())
	}
}
