// @Time : 2018/6/9 11:08
// @Author : minigeek
package main

import (
	"mini_crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/search", controller.CreateSearchResultHandler(
		"mini_crawler/frontend/view/search.html",
	))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}