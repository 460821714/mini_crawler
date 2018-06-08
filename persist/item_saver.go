// @Time : 2018/6/7 14:00
// @Author : minigeek
package persist

import (
	"log"

	"context"

	"crawler/engine"

	"github.com/pkg/errors"
	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	go func() {
		var itemCount int
		for {
			item := <-out
			log.Printf("ItemSaver got item %d,%v", itemCount, item)
			itemCount++
			err := save(client, index, item)
			if err != nil {
				log.Printf("item: %v,save faild:%v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("Must supply type")
	}
	indexService := client.Index().
		Index(index).
		Type(item.Type)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.BodyJson(item).
		Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
