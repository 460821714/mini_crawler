// @Time : 2018/6/8 16:36
// @Author : minigeek
package engine

import (
	"log"
	"mini_crawler/fetcher"
)

func Worker(request Request) (ParseResult, error) {
	//log.Printf("Fetching %s", request.Url)
	contents, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("%v", err)
		return ParseResult{}, err
	}
	parseResult := request.Parser.Parse(contents, request.Url)
	return parseResult, nil
}
