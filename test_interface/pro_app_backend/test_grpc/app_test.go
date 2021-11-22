package test_grpc

import (
	"context"
	"fmt"
	v1 "newgitlab.com/xg-pro/appbackend/api/appbackend/v1"
	"testing"
)

var email string = "testUser@daily_practice.com"

func Test_Register(t *testing.T) {
	client := NewAppClient()
	reply, err := client.Register(context.Background(), &v1.RegisterRequest{
		Email:    email,
		Password: "a123456",
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", reply)
}