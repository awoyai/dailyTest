package test_grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/dipperin/go-ms-toolkit/json"
	v1 "newgitlab.com/xg-pro/usercenter/api/usercenter/v1"
)

var backEmail string = "superAdmin@yyyz.com"

func Test_BackendLogin(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.BackendLogin(context.Background(), &v1.LoginReq{
		Identifier:   backEmail,
		Certificate:  "F9MDPmSov62chnOV",
		IdentityType: v1.IdentityType_Password,
		Need:         true,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("token: " + reply.Token)
	fmt.Printf("roleId: %v \n", reply.RoleId)
	json := json.StringifyJson(reply.List)
	fmt.Printf("role: %v", json)
}

func Test_BackendAuth(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.BackendAuth(context.Background(), &v1.BackendAuthReq{
		UserId: 1,
		Token:  "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzY5NjQ4NDMsInJpZCI6MSwidWlkIjoxfQ.kYbUyBotwuJnrzcoMOtMwRcF82GeHbUXc0iqZaQC_fIGCD7Zw480t1hpbaAjtyHgmjLeBmt3tooSsALxjYptl4qdTcUtwbZReJQtNMaYeaR3ABKALAFPIfY82BcrLAxlRWA1E3LkOqWUR19BEiAcwWIn-rGVqfRP7cgAz3T2EIkgF6YX0LxiCg10vI8mpjGL8pUdkwItZ245G3HO-Fkj7llb_BG1XNO5phZbNdNoTEydPDHGw9aDc9ISLzCn4LO27zMFqVKrsWwmogyzbiKcPUeUMLrRTMG8-3GNy0QvqsVsfEXjjVQJMB-5SQIQNrAIpd6BUZr4T9k1niq1OhB12g",
		Path:   "daily_practice",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("reply: %v", json.StringifyJson(reply))
}

func Test_CreateResource(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.CreateResource(context.Background(), &v1.ResourceReq{
		Path:         "test_path",
		Tittle:       "test_tittle",
		ParentId:     0,
		Description:  "daily_practice",
		ResourceType: v1.ResourceType_Menu,
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_CreateRole(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.CreateRole(context.Background(), &v1.RoleReq{
		Name:        "Admin",
		Description: "财务",
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_BindRoleResource(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.BindRoleResource(context.Background(), &v1.BindRoleResourceReq{
		RoleId:       1,
		ResourceList: []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9},
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_ListUser(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.ListUser(context.Background(), &v1.ListUserReq{
		Page:      1,
		PageSize:  2,
		Condition: "0",
		RoleId:    2,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("users: %v", json.StringifyJson(reply))
}

func Test_UserDetails(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.UserDetails(context.Background(), &v1.UserId{UserId: 5})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("users: %v", json.StringifyJson(reply))
}

func Test_OperateUser(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.OperateUser(context.Background(), &v1.OperateUserReq{
		UserId:      7,
		AdminId:     2,
		OperateType: v1.OperateType_Freeze,
		Status:      true,
		Remark:      "daily_practice freeze",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("reply: %v", json.StringifyJson(reply))
}

func Test_SetUpPwd(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.SetUpPwd(context.Background(), &v1.SetUpPwdReq{
		UserId:       5,
		Identifier:   email,
		Certificate:  "a123456",
		IdentityType: v1.IdentityType_PayPassword,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("reply: %v", json.StringifyJson(reply))
}
