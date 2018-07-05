// @Time : 2018/5/23 20:19
// @Author : minigeek
package engine

type ParseFunc func([]byte, string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// request body.
type Request struct {
	// request url
	Url string
	// parser
	Parser Parser
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

type NilParserFunc struct{}

func (NilParserFunc) Parse(contents []byte, url string) ParseResult {
	return ParseResult{}
}

func (NilParserFunc) Serialize() (name string, args interface{}) {
	return "ParseNil", nil
}

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(f ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: f,
		name:   name,
	}
}
