package pro_news

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "newgitlab.com/xg-pro/news/api/news/v1"
)

func NewNewsClient() v1.NewsClient {
	//conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("10.8.0.16:34620"))
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:30257"))
	if err != nil {
		panic(err)
	}
	return v1.NewNewsClient(conn)
}