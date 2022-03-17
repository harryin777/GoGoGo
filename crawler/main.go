/**
* @Author: yky
* @Date : 2021/11/17 9:16 下午
* @Descriptions :
 */
package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36"),
	)

	c.SetRequestTimeout(60 * time.Second)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//	//每次使用不同的 userAgent, 防止服务器对同一个 userAgent 做限制
	//	r.Headers.Set("User-Agent", RandomString())
	//})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting:", request.URL)
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 5,
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.Visit("https://www.zhihu.com/")
}

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
