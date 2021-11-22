package test_grpc

import (
	"context"
	"fmt"
	v1 "newgitlab.com/xg-pro/usercenter/api/usercenter/v1"
	"testing"
)

func Test_send_vcode(t *testing.T) {
	client := NewUsercenterClient()

	for i := 0; i < 50; i++ {
		go func() {
			_, err := client.SendEmailVCode(context.Background(), &v1.VCodeReq{
				Email:     fmt.Sprintf("1174734104@qq.com"),
				ApplyType: v1.ApplyType_Register,
				SendType:  v1.SendType_Email,
			})
			fmt.Println("err", err)
		}()
	}
	for {
		fmt.Scanln()
	}
}
