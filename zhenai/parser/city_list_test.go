// @Time : 2018/5/24 11:29
// @Author : minigeek
package parser

import (
	"crawler/fetcher"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city_list_test_data.html")
	if err != nil {
		fmt.Println(err)
	}
	result := ParseCityList(contents, "")
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests;but had %d", resultSize, len(result.Requests))
	}

	exceptedUrls := []string{"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng"}

	for i, r := range exceptedUrls {
		if result.Requests[i].Url != r {
			t.Errorf("excepted url is %s;but was %s", r, result.Requests[i].Url)
		}
	}
}

func TestFetch(t *testing.T) {

	contents, err := fetcher.Fetch("http://album.zhenai.com/u/1314495053")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(contents))
}
