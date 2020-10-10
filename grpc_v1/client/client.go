package main

import (
	"context"
	"fmt"
	"mathproto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 创建一个 grpc 连接，Dial() 方法直连 gRPC 服务
	conn, err := grpc.Dial("127.0.0.1:5555", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	// 创建一个专门调用 math 服务的客户端
	client := mathproto.NewMathClient(conn)
	// 创建一个请求超时时间为1秒的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 调用 Plus 服务
	r1, err := client.Plus(ctx, &mathproto.PlusRequest{A: 1, B: 2})
	if err != nil {
		panic(err.Error())
	}
	r2, err := client.Mult(ctx, &mathproto.MultRequest{A: 3, B: 4})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r1.C, r2.C)
}
