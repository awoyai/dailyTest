package test_grpc

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "newgitlab.com/xg-pro/usercenter/api/usercenter/v1"
)

func NewUsercenterClient() v1.UsercenterClient {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("0.0.0.0:30245"))
	//conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("10.8.0.18:27830"))
	if err != nil {
		panic(err)
	}
	return v1.NewUsercenterClient(conn)
}