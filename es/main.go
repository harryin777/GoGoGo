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
	DaysRetentionRate()

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

func DaysRetentionRate() {
	var retentionDays = 4
	startDate := "2022-05-12 00:00:00"
	field := "did.keyword"
	indexName := "oppo_api"

	d1, _ := time.Parse(Utils.FullLayout, startDate)
	d2 := d1.Add(24 * time.Hour)
	d1Count := 0
	uidMap := make(map[string]int)

	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetURL(cfg.URL), elastic.SetBasicAuth(cfg.Username, cfg.Password), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
	col := elastic.NewCollapseBuilder(field)
	for i := 0; i < retentionDays; i++ {
		rq := elastic.NewRangeQuery("time").Gt(d1.Format(Utils.FullLayout)).Lte(d2.Format(Utils.FullLayout))
		searchResult, err := client.Search(indexName).Query(rq).Collapse(col).Pretty(true).Size(max).Do(ctx)
		if err != nil {
			panic(err)
		}
		for index, hit := range searchResult.Hits.Hits {
			uids, found := hit.Fields.Strings(field)
			if !found {
				fmt.Printf("hits %v can`t find %v \n", index, field)
				continue
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

		d1 = d1.Add(time.Hour * 24)
		d2 = d2.Add(time.Hour * 24)
	}

	finalCount := 0
	for _, val := range uidMap {
		if val == retentionDays {
			finalCount++
		}
	}

	if d1Count == 0 {
		fmt.Println("d1Count is zero")
	} else {
		fmt.Printf("%v days` retention rate: %v \n", retentionDays, fmt.Sprintf("%.2f", float64(finalCount)/float64(d1Count)))
	}

}
