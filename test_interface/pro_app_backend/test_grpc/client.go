package test_grpc

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "newgitlab.com/xg-pro/appbackend/api/appbackend/v1"
)

func NewAppClient() v1.UserClient {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("10.8.0.16:9000"))
	if err != nil {
		panic(err)
	}
	return v1.NewUserClient(conn)
}