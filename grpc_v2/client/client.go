package main

import (
	"context"
	"fmt"
	"mathproto"
	"time"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
)

func main() {
	for i := 0; i < 10; i++ {
		address := findServiceByName("MathService_v1")
		calc(address)
		time.Sleep(time.Second)
	}
}

func findServiceByName(name string) string {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	res, err := etcdClient.Get(ctx, name)
	if err != nil {
		panic(err.Error())
	}
	return string(res.Kvs[0].Value)
}

func calc(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	gRPCClient := mathproto.NewMathClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	res, err := gRPCClient.Plus(ctx, &mathproto.PlusRequest{A: 5, B: 6})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res.C)
}
