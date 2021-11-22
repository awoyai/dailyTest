package test_grpc

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
	v1 "newgitlab.com/xg-pro/usercenter/api/usercenter/v1"
	"newgitlab.com/xg-pro/usercenter/pkg/types"
)

var email string = "1174734104@qq.com"

func Test_QueryUserExist(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.QueryUserExist(context.Background(), &v1.EmailReq{Email: email})
	if err != nil {
		t.Error(err)
	}
}

func Test_Register(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.Register(context.Background(), &v1.RegisterReq{
		Email:    "1111@qq.com",
		Password: "a123456",
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_ResetPwd(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.ResetPwd(context.Background(), &v1.ResetPwdReq{
		Identifier:   email,
		NewPassword:  "a123456",
		IdentityType: v1.IdentityType_Password,
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_Login(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.Login(context.Background(), &v1.LoginReq{
		Identifier:   email,
		Certificate:  "a123456",
		IdentityType: v1.IdentityType_Password,
		Need:         true,
		CaptchaReq: &v1.CaptchaReq{
			Ticket:  "t03HqCtslTMGKrpQ09TMZe-dHrAxuC3Fjczbfod5KHRtCEA2udg4G9tLGnULqFjBjFAQb4-gqOgFZKSPJGAHnQsC-5ISulvUgdywYTHVBUCdB-Iz4p3w8N2ww**",
			Randstr: "@KeV",
		},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("token: " + reply.Token)
}

func Test_AuthToken(t *testing.T) {
	client := NewUsercenterClient()
	reply, err := client.Login(context.Background(), &v1.LoginReq{
		Identifier:   email,
		Certificate:  "a123456",
		IdentityType: v1.IdentityType_Password,
		Need:         true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("token:%s, roleId: %v \n", reply.Token, reply.RoleId)
	_, err = client.AuthToken(context.Background(), &v1.AuthTokenReq{
		UserId: reply.UserId,
		Token:  reply.Token,
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_SendEmailVCode(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.SendEmailVCode(context.Background(), &v1.VCodeReq{
		Email:     email,
		ApplyType: v1.ApplyType_Login,
		SendType:  v1.SendType_Email,
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_AuthVCode(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.AuthVCode(context.Background(), &v1.AuthVCodeReq{
		Email:     email,
		ApplyType: v1.ApplyType_Login,
		SendType:  v1.SendType_Email,
		Code:      "163681",
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_GenToken(t *testing.T) {
	client := NewUsercenterClient()
	token, err := client.GenToken(context.Background(), &v1.UserId{
		UserId: 5,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println("token: " + token.Token)
}

func Test_AuthSingleToken(t *testing.T) {
	client := NewUsercenterClient()
	_, err := client.AuthToken(context.Background(), &v1.AuthTokenReq{
		UserId: 7,
		Token:  "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzU4MzYzNDYsInJpZCI6MiwidWlkIjo3fQ.ZmuvHLKhL0XXtV2biIg6Bfob1daZu_j561g38O2mgezXNtgkJ13w_H4RD71U5T0qZavSZ9HcYfTzyqXkNHMdI7gNPabvfHZ7nBOL0DYhWfkIx3tQeKp9n35wSAAgUKLtnTh3x-Wyb4U-VKqSI33Cl8dZfKqCz8NU_n6Q1E_4-XKTAwfAg-Tla2KZTjOQDfABmAeiseseeM7HE4_Fo-lwyVRKV2tubydKie0zEAHq_lx4MZbRqc7pqq0yb0_r_v-2m1tQhEfFp1FqWDGEYGZfyZ8-6bkxThtzrUyW2dSUeylbzPrRCd-GsdnseXbPCzzal_1FFXHsUgn4bxVgye0TEw",
	})
	if err != nil {
		t.Error(err)
	}
}

func Test_CompanyUser(t *testing.T) {
	client := NewUsercenterClient()
	user, err := client.CompanyUser(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user.UserId)
}

func Test_QueryUser(t *testing.T) {
	client := NewUsercenterClient()
	user, err := client.UserDetails(context.Background(), &v1.UserId{UserId: 1})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", user)
}

func Test_QueryStatus(t *testing.T) {
	client := NewUsercenterClient()
	status, err := client.QueryStatus(context.Background(), &v1.UserId{UserId: 1})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", status)
}

func Test_GenResourceList(t *testing.T) {
	resources := []*types.Resource{
		{
			Model:       gorm.Model{ID: 1},
			ParentId:    0,
			Name:        "11",
			Resource:    "11",
			Description: "11",
		}, {
			Model:       gorm.Model{ID: 2},
			ParentId:    1,
			Name:        "11",
			Resource:    "11",
			Description: "11",
		}, {
			Model:       gorm.Model{ID: 3},
			ParentId:    1,
			Name:        "11",
			Resource:    "11",
			Description: "11",
		}, {
			Model:       gorm.Model{ID: 4},
			ParentId:    0,
			Name:        "11",
			Resource:    "11",
			Description: "11",
		}, {
			Model:       gorm.Model{ID: 5},
			ParentId:    4,
			Name:        "11",
			Resource:    "11",
			Description: "11",
		}, {
			Model:       gorm.Model{ID: 6},
			ParentId:    2,
			Name:        "11",
			Resource:    "11",
			Description: "11",
		}, {
			Model:       gorm.Model{ID: 7},
			ParentId:    6,
			Name:        "11",
			Resource:    "11",
			Description: "11",
		},
	}
	pMap := make(map[uint64][]*Resource)
	rMap := make(map[uint64]*Resource)
	for _, v := range resources {
		resource := &Resource{
			ResourceId: uint64(v.ID),
			Path:       v.Resource,
			Tittle:     v.Name,
		}
		rMap[uint64(v.ID)] = resource
		pMap[uint64(v.ParentId)] = append(pMap[uint64(v.ParentId)], resource)
	}

	if _, ok := pMap[0]; !ok {
		return
	}

	for k, v := range pMap {
		if _, ok := rMap[k]; !ok {
			continue
		}
		rMap[k].Children = v
	}
	fmt.Printf("%+v", pMap[0])
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId uint64      `protobuf:"varint,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Path       string      `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Tittle     string      `protobuf:"bytes,3,opt,name=tittle,proto3" json:"tittle,omitempty"`
	Children   []*Resource `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
}

func Test_CreatPayPassword(t *testing.T) {
	client := NewUsercenterClient()
	status, err := client.SetUpPwd(context.Background(), &v1.SetUpPwdReq{
		UserId:       4,
		Identifier:   email,
		Certificate:  "123456",
		IdentityType: v1.IdentityType_PayPassword,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", status)
}
