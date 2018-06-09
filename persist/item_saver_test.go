// @Time : 2018/6/8 9:57
// @Author : minigeek
package persist

import (
	"context"
	"encoding/json"
	"mini_crawler/engine"
	"mini_crawler/model"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	item := engine.Item{
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

	// try to connect to elasticsearch client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "crawler_test"
	// save item
	err = save(client, index, item)
	if err != nil {
		panic(err)
	}

	// fetch item
	result, err := client.Get().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	var actual engine.Item
	err = json.Unmarshal(*result.Source, &actual)
	if err != nil {
		panic(err)
	}
	actualProfile, err := model.FromJsonObj(actual.Payload)
	if err != nil {
		panic(err)
	}
	actual.Payload = actualProfile

	// verfiy item
	if item != actual {
		t.Errorf("excepted %v,but is %v", item, actual)
	}
}
