// @Time : 2018/6/8 19:50
// @Author : minigeek
package model

import "crawler/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
