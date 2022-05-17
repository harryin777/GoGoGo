package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/olivere/elastic/v7/config"
	"sync"
	"test1/Utils"
	"test1/common/loader"
	"time"
)

var (
	once sync.Once
	cfg  *config.Config
)

//const max = 100000000
const max = 100

func init() {
	once.Do(func() {
		loader.LoadElasticSearchConfig(&cfg, "elasticsearch.json")
	})
}

func main() {
	FiveDayRetentionRate()

}

func initial() {
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
}

func FiveDayRetentionRate() {
	var retentionDays = 5
	startDate := "2022-05-10 00:00:00"
	d1, _ := time.Parse(Utils.FullLayout, startDate)
	d1Count := 0
	uidMap := make(map[string]int)

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(cfg.URL), elastic.SetBasicAuth(cfg.Username, cfg.Password), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
	col := elastic.NewCollapseBuilder("uid.keyword")
	for i := 0; i < retentionDays; i++ {
		d2 := d1.Add(time.Hour * 24)
		rq := elastic.NewRangeQuery("time").Gt().Lte(d2)
		searchResult, err := client.Search("oppo_api").Query(rq).Collapse(col).Pretty(true).Size(max).Do(ctx)
		if err != nil {
			panic(err)
		}
		for _, hit := range searchResult.Hits.Hits {
			uids, found := hit.Fields.Strings("uid.keyword")
			if !found {
				fmt.Println("can`t find uid.keyword")
				return
			}
			if len(uids) != 0 {
				if val, exists := uidMap[uids[0]]; exists {
					uidMap[uids[0]] = val + 1
				} else {
					uidMap[uids[0]] = 1
				}

			}
		}
		if i == 0 {
			d1Count = len(uidMap)
		}

		d1 = d2.Add(time.Hour * 24)
	}

	finalCount := 0
	for _, val := range uidMap {
		if val == retentionDays {
			finalCount++
		}
	}

	fmt.Printf("%v retention rate: %v \n", retentionDays, finalCount/d1Count)

}
