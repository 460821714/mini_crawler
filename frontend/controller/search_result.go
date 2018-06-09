// @Time : 2018/6/9 11:02
// @Author : minigeek
package controller

import (
	"mini_crawler/frontend/view"

	"net/http"

	"strconv"

	"strings"

	"mini_crawler/frontend/model"

	"context"

	"mini_crawler/engine"
	"reflect"

	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	condition := strings.TrimSpace(req.FormValue("q"))
	page, err := s.getSearchResult(condition, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = s.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s SearchResultHandler) getSearchResult(condition string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	res, err := s.client.Search("crawler").
		Query(elastic.NewQueryStringQuery(condition)).
		From(from).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = res.TotalHits()
	result.Start = from
	result.Items = res.Each(reflect.TypeOf(engine.Item{}))
	return result, nil
}
