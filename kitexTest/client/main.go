package main

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"log"
	snow2 "test1/kitexTest/kitex_gen/snow"
	"test1/kitexTest/kitex_gen/snow/snow"
	"time"
)

func main() {
	client, err := snow.NewClient("snow", client2.WithHostPorts("0.0.0.0:8888"))
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
