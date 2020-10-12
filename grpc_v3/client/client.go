package main

import (
	"context"
	"fmt"
	"lb"
	"mathproto"
	"time"

	"google.golang.org/grpc/resolver"

	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {
	serviceName := "MathService"
	grpc.UseCompressor(gzip.Name)
	addressList := getServerAddressList(serviceName)
	resolver.Register(&lb.MyResolverBuilder{
		ServiceName: serviceName,
		AddressList: addressList,
	})
	for i := 0; i < 5; i++ {
		callServer(serviceName)
	}
}

func getServerAddressList(name string) []string {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := etcdClient.Get(ctx, name, clientv3.WithPrefix())
	if err != nil {
		panic(err.Error())
	}
	addressList := make([]string, len(res.Kvs))
	for i, val := range res.Kvs {
		addressList[i] = string(val.Value)
	}
	return addressList
}

func callServer(target string) {
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	conn, err := grpc.DialContext(
		ctx1,
		target,
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy":"%s"}`, roundrobin.Name)),
	)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	gRPCClient := mathproto.NewMathClient(conn)
	res, _ := gRPCClient.Plus(ctx2, &mathproto.PlusRequest{A: 8, B: 2})
	fmt.Println(res.C)
}
