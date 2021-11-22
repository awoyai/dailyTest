package pro_news

import (
	"context"
	"fmt"
	"testing"

	v1 "newgitlab.com/xg-pro/news/api/news/v1"
)

func Test_CreateNews(t *testing.T) {
	client := NewNewsClient()
	reply, err := client.CreateNews(context.Background(), &v1.NewsReq{
		Author:   "testAuthor",
		Title:    "testTitle",
		TagIds:   []uint64{1, 2},
		CoverImg: "test_img_url",
		Article:  "test_article",
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", reply)
}

func Test_ListNews(t *testing.T) {
	client := NewNewsClient()
	news, err := client.ListNews(context.Background(), &v1.ListNewsReq{
		Page:     0,
		PageSize: 0,
		TagId:    1,
		VipOnly:  1,
		TopPicks: 4,
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", news)
}

func Test_OperateNews(t *testing.T) {
	client := NewNewsClient()
	_, err := client.OperateNews(context.Background(), &v1.OperateReq{
		NewsId:  11111,
		Operate: v1.OperateType_Publish,
		Status:  true,
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}

func Test_ListRecommend(t *testing.T) {
	client := NewNewsClient()
	reply, err := client.ListRecommend(context.Background(), &v1.ListRecommendReq{
		TagIds: []uint64{1, 2},
		Limit:  3,
		NewsId: 1,
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("reply: %v", reply)
}

func Test_CreateTag(t *testing.T) {
	client := NewNewsClient()
	_, err := client.CreateTag(context.Background(), &v1.CreateTagReq{
		Tag:         "Starsharks",
		Description: "Starsharks",
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}

func Test_QueryDetails(t *testing.T) {
	client := NewNewsClient()
	details, err := client.QueryNewsDetails(context.Background(), &v1.NewsID{
		NewsId: 1})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("details: %v", details)
}

func Test_UpdateNews(t *testing.T) {
	client := NewNewsClient()
	_, err := client.UpdateNews(context.Background(), &v1.UpdateNewsReq{
		NewsId:   1,
		Author:   "updateAuthor",
		Title:    "updateTitle",
		TagIds:   []uint64{5, 6},
		CoverImg: "updateCoverImg",
		Article:  "111",
	})
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}
