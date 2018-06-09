// @Time : 2018/5/28 15:39
// @Author : minigeek
package parser

import (
	"fmt"
	"io/ioutil"
	"mini_crawler/engine"
	"mini_crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {

	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		fmt.Println(err)
	}
	result := ParseProfile(contents, "http://album.zhenai.com/u/1314495053", "风中的蒲公英")
	if len(result.Items) != 1 {
		t.Errorf("Items length should be 1;but was %d", len(result.Items))
	}
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "profile",
		Id:   "1314495053",
		Payload: model.Profile{
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
	actual := result.Items[0]
	if expected != actual {
		t.Errorf("Item should be %v;but was %v", expected, actual)
	}
}
