package main

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/retry"
	"log"
	snow2 "test1/kitexTest/kitex_gen/snow"
	"test1/kitexTest/kitex_gen/snow/snow"
	"time"
)

func main() {
	fp := failureRetryPolicy()
	//配置RPC连接超时
	rpcTimeout := client2.WithRPCTimeout(3 * time.Second)
	//配置连接超时
	connTimeout := client2.WithConnectTimeout(50 * time.Millisecond)
	client, err := snow.NewClient("snow",
		client2.WithHostPorts("0.0.0.0:8888"),
		client2.WithLongConnection(connpool.IdleConfig{10, 1000, time.Minute}),
		rpcTimeout,
		connTimeout,
		client2.WithFailureRetry(fp))
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	wantedReq := &snow2.SnowRequest{Wanted: 6}
	wantedRes, err := client.Wanted(context.Background(), wantedReq)
	if err != nil {
		panic(err)
	}
	log.Println(wantedRes)
}

func failureRetryPolicy() *retry.FailurePolicy {
	fp := retry.NewFailurePolicy()
	fp.WithMaxRetryTimes(3) // 配置最多重试3次

	// 总耗时，包括首次失败请求和重试请求耗时达到了限制的duration，则停止后续的重试。
	fp.WithMaxDurationMS(100)

	// 关闭链路中止
	fp.DisableChainRetryStop()

	// 开启DDL中止
	fp.WithDDLStop()

	// 退避策略，默认无退避策略
	fp.WithFixedBackOff(10)     // 固定时长退避
	fp.WithRandomBackOff(5, 10) // 随机时长退避

	// 开启重试熔断
	fp.WithRetryBreaker(0.5)

	// 同一节点重试
	fp.WithRetrySameNode()
	return fp
}
