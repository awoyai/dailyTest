package pro_invitation

import (
	"context"
	"fmt"
	"testing"

	v1 "newgitlab.com/xg-pro/invitation/api/invitation/v1"
)

func Test_GenInviteCode(t *testing.T) {
	client := NewInvitationClient()
	reply, err := client.GenInviteCode(context.Background(), &v1.GenInviteCodeReq{
		UserId: 1,
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", reply)
}

func Test_CheckInviteCode(t *testing.T) {
	client := NewInvitationClient()
	reply, err := client.CheckInviteCode(context.Background(), &v1.CheckInviteCodeReq{
		InviteCode: "XhYSwE",
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", reply)
}

func Test_CreateInvitedRecord(t *testing.T) {
	client := NewInvitationClient()
	reply, err := client.CreateInvitedRecord(context.Background(), &v1.CreateInvitedRecordReq{
		UserId:        1,
		InvitedUserId: 2,
		Code:          "XhYSwE",
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", reply)
}

func Test_ListFans(t *testing.T) {
	client := NewInvitationClient()
	reply, err := client.ListFansByUserId(context.Background(), &v1.ListFansByUserIdReq{
		UserId:   1,
		Page:     1,
		PerPage:  10,
		Paginate: true,
		OrderBy:  "created_at desc",
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", reply)
}
