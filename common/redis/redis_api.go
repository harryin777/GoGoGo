/**
*   @Author: yky
*   @File: redis_api
*   @Version: 1.0
*   @Date: 2021-05-13 22:18
 */
package redis

import (
	"fmt"
	"test1/common/loader"
	"time"
)

package redis

import (
"allmusic-api-server/common/loader"
"context"
"fmt"
"github.com/go-redis/redis/v8"
"time"
)

type Instance struct {
	MasterClient redis.UniversalClient
}

type datasource struct {
	Addr        string
	Password    string
	DB          int
	ClusterMode bool
	PoolSize    int
}

const redisConfigFile = "redis.cfg"

var redis *Instance

func init() {
	redis = loadCfgDataSource("main")
}

func MainRedis() redis.UniversalClient {
	return redis.MasterClient
}

func loadCfgDataSource(section string) *Instance {
	cfg := loader.LoadConfigFile(redisConfigFile)
	addr, err := cfg.GetString(section, "Addr")
	if err != nil {
		msg := fmt.Sprintf("Read redis config error. miss Addr. section: %v err: %v", section, err)
		panic(msg)
	}
	password, err := cfg.GetString(section, "Password")
	if err != nil {
		msg := fmt.Sprintf("Read redis config error. miss Password. section: %v err: %v", section, err)
		panic(msg)
	}
	db, err := cfg.GetInt64(section, "DB")
	if err != nil {
		msg := fmt.Sprintf("Read redis config error. miss DB. section: %v err: %v", section, err)
		panic(msg)
	}
	maxClients, _ := cfg.GetInt64(section, "PoolSize")
	clusterMode, _ := cfg.GetBool(section, "ClusterMode")
	return createDataSource(&datasource{
		Addr:        addr,
		Password:    password,
		DB:          int(db),
		ClusterMode: clusterMode,
		PoolSize:    int(maxClients),
	})
}

func createDataSource(info *datasource) *Instance {
	dataSource := &Instance{}
	if info.ClusterMode {
		clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:          []string{info.Addr},
			DialTimeout:    time.Second * 10,
			IdleTimeout:    time.Second * 5,
			PoolSize:       info.PoolSize,
			RouteByLatency: true,
		})

		status := clusterClient.Ping(context.Background()).Val()
		if status != "PONG" {
			msg := fmt.Sprintf("can not ping redis. info: %+v. status: %v", info, status)
			panic(msg)
		}

		dataSource.MasterClient = clusterClient
	} else {
		cfg := &redis.Options{}
		cfg.Addr = info.Addr
		cfg.Password = info.Password
		cfg.DB = info.DB
		cfg.DialTimeout = time.Second * 10
		cfg.IdleTimeout = time.Second * 5
		if info.PoolSize != 0 {
			cfg.PoolSize = info.PoolSize
		}

		client := redis.NewClient(cfg)
		status := client.Ping(context.Background()).Val()
		if status != "PONG" {
			msg := fmt.Sprintf("can not ping redis. cfg: %+v. status: %v", cfg, status)
			panic(msg)
		}
		dataSource.MasterClient = client
	}

	return dataSource
}
