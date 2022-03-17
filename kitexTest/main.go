package main

import (
	"log"
	"test1/kitexTest/kitex_gen/snow/snow"
)

func main() {
	svr := snow.NewServer(new(SnowImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
