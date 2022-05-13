package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"github.com/ysmood/got/lib/gop"
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

	//client, err := elastic.NewClientFromConfig(cfg)
	//sniff 集群导致请求超时
	//client, err := elastic.NewClient(elastic.SetURL(cfg.URL), elastic.SetBasicAuth(cfg.Username, cfg.Password), elastic.SetSnifferTimeout(1*time.Nanosecond))
	//停止连接集群时嗅探
	client, err := elastic.NewClient(elastic.SetURL(cfg.URL), elastic.SetBasicAuth(cfg.Username, cfg.Password), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println("connected to elasticsearch")
	//检查索引存在
	exists, err := client.IndexExists("oppo_info").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Println(exists)

	result, err := client.Refresh("oppo_info").Do(ctx)
	if err != nil {
		panic(err)
	}
	gop.P(result)
}
