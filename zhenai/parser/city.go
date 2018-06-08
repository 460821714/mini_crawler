// @Time : 2018/5/26 16:38
// @Author : minigeek
package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	cityReg    = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]*)</a>`) // regexp of city parse
	cityUrlReg = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

// ParseCity returns engine.ParseResult after parse city  with contents.
func ParseCity(contens []byte, _ string) engine.ParseResult {
	items := cityReg.FindAllSubmatch(contens, -1)
	result := engine.ParseResult{}
	for _, item := range items {
		name := string(item[2])
		//result.Item = append(result.Item, "Name:"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParseFunc: ProfileParse(name),
		})
	}

	items = cityUrlReg.FindAllSubmatch(contens, -1)
	for _, v := range items {
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(v[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
