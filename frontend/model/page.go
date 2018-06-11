// @Time : 2018/6/8 19:50
// @Author : minigeek
package model

type SearchResult struct {
	Hits     int64
	Start    int
	PrevFrom int
	NextFrom int
	Query    string
	Items    []interface{}
}
