// @Time : 2018/6/8 19:59
// @Author : minigeek
package view

import (
	"mini_crawler/engine"
	"mini_crawler/frontend/model"
	common "mini_crawler/model"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("search.html")
	out, err := os.Create("search_test.html")
	if err != nil {
		panic(err)
	}
	page := model.SearchResult{}
	page.Hits = 100
	page.Start = 0
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "profile",
		Id:   "1314495053",
		Payload: common.Profile{
			Name:       "风中的蒲公英",
			Gender:     "女",
			Age:        41,
			Height:     158,
			Weight:     48,
			Income:     "3001-5000元",
			Marriage:   "离异",
			Education:  "中专",
			Occupation: "公务员",
			Hokou:      "四川阿坝",
			Xinzuo:     "处女座",
			House:      "已购房",
			Car:        "未购车",
		},
	}
	for i := 0; i < 100; i++ {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
