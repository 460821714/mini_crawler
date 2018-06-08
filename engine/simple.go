// @Time : 2018/5/24 10:48
// @Author : minigeek
package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	taskQueue := seeds
	for len(taskQueue) > 0 {
		request := taskQueue[0]
		taskQueue = taskQueue[1:]

		//fetch
		parseResult, err := Worker(request)
		if err != nil {
			continue
		}
		taskQueue = append(taskQueue, parseResult.Requests...)

		//print item.
		for _, v := range parseResult.Items {
			log.Println(v)
		}
	}
}
