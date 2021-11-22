package pro_invitation

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "newgitlab.com/xg-pro/invitation/api/invitation/v1"
)

func NewInvitationClient() v1.InvitationClient {
	//conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("10.8.0.16:34620"))
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		panic(err)
	}
	return v1.NewInvitationClient(conn)
}
