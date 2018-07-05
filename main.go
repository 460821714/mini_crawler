// @Time : 2018/5/23 15:07
// @Author : minigeek
package main

import (
	"log"
	"mini_crawler/engine"
	"mini_crawler/persist"
	"mini_crawler/scheduler"
	"mini_crawler/zhenai/parser"
)

// start url for fetch.
const startUrl = "http://www.zhenai.com/zhenghun"

func main() {
	itemChan, err := persist.ItemSaver("crawler")
	if err != nil {
		panic(err)
	}
	log.Println("start fetch...")
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:    startUrl,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})

	//e.Run(engine.Request{
	//	Url:       "http://www.zhenai.com/zhenghun/shanghai",
	//	ParseFunc: parser.ParseCity,
	//})
}
