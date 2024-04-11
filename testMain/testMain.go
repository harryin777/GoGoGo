package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		fmt.Printf("here : %v \n", i)
		ch <- vs[i](name)
	}
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

//func main() {
//	ret := exec("111", func(n string) string {
//		return n + "func0"
//	}, func(n string) string {
//		return n + "func1"
//	}, func(n string) string {
//		return n + "func2"
//	}, func(n string) string {
//		return n + "func3"
//	})
//	_ = ret
//	fmt.Println(ret)
//
//	var s SSS
//	s.impl = NewImpl()
//	s.impl.Say("111")
//
//}

type Data struct {
	num int
}

type Impl2 struct {
	a Data
}

func main() {
	//i := &Impl2{
	//	a: Data{num: 10},
	//}
	//
	//i.Value()
	//
	//fmt.Println(i.a.num) // 输出的是0

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "testset",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", "testst")
}

func (i Impl2) Value() {
	i.a = Data{num: 0} // 重新初始化了a字段
}
