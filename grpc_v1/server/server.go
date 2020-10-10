package main

import (
	"context"
	"mathproto"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type Math struct{}

func (m *Math) Plus(ctx context.Context, request *mathproto.PlusRequest) (*mathproto.PlusResponse, error) {
	return &mathproto.PlusResponse{C: request.A + request.B}, nil
}

func (m *Math) Mult(ctx context.Context, request *mathproto.MultRequest) (*mathproto.MultResponse, error) {
	return &mathproto.MultResponse{C: request.A * request.B}, nil
}

func main() {
	// 创建一个 TCP 监听服务
	lis, err := net.Listen("tcp", "127.0.0.1:5555")
	if err != nil {
		panic(err.Error())
	}
	// 创建一个 grpc 服务
	s := grpc.NewServer()
	// grpc 服务映射到结构体 Math
	mathproto.RegisterMathServer(s, &Math{})
	// 通过反射解析服务
	reflection.Register(s)
	// 将 grpc 服务绑定到 TCP 监听服务上
	if err := s.Serve(lis); err != nil {
		panic(err.Error())
	}
}
