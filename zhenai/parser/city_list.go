// @Time : 2018/5/23 20:23
// @Author : minigeek
package parser

import (
	"mini_crawler/engine"
	"regexp"
)

const cityListRegular = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]*)</a>` // regexp of city list parse

// ParseCityList returns engine.ParseResult after parse city list with contents.
func ParseCityList(contents []byte, _ string) engine.ParseResult {
	reg := regexp.MustCompile(cityListRegular)
	matchers := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range matchers {
		//result.Item = append(result.Item, "City:"+string(v[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:    string(v[1]),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity")})
	}
	return result
}
