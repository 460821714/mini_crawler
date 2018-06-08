// @Time : 2018/5/23 20:19
// @Author : minigeek
package engine

// request body.
type Request struct {
	// request url
	Url string
	// parser
	ParseFunc func([]byte, string) ParseResult
}

// this is parse result.
type ParseResult struct {
	// requests that need to go to parse after parse request
	Requests []Request
	// data item after parse request
	Items []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParserFunc(b []byte) ParseResult {
	return ParseResult{}
}
