package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"sync"
	"test1/common/loader"
)

var (
	once sync.Once
	cfg  *config.Config
)

func init() {
	once.Do(func() {
		loader.LoadElasticSearchConfig(&cfg, "elasticsearch.json")
	})
}

func main() {
	ctx := context.Background()

	client, err := elastic.NewClientFromConfig(cfg)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connected to elasticsearch")
	exists, err := client.IndexExists("oppo_info").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println(exists)

}
