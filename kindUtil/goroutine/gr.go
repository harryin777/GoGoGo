package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	batchSize := 30
	batchCount := 10

	eg := errgroup.Group{}

	c := make(chan int)
	d := make(chan int)
	d2 := make(chan int)
	defer close(c)
	defer close(d)

	for i := 0; i < batchSize; i++ {
		if i > 0 && i%batchCount == 0 {
			tmp := i
			// 批次开始执行
			eg.Go(func() error {
				if res, err := job(tmp); err != nil {
					c <- res
					return err
				}
				return nil
			})
		}
	}

	go func() {
		for {
			select {
			case data := <-c:
				rollback()
				fmt.Println(data)
				d2 <- 1
				close(d2)
			case <-d:
				fmt.Println("结束")
				close(d2)
				return
			}
		}
	}()

	if err := eg.Wait(); err == nil {
		d <- 1
	}

	<-d2
	fmt.Println("真的结束")
}

func job(batch int) (res int, err error) {
	fmt.Printf("Processing batch start %d\n", batch)
	// 模拟错误
	if batch == 10 {
		return res, fmt.Errorf("simulated error in batch %d", batch)
	}

	fmt.Printf("Processing batch end %d\n", batch)

	return 0, nil
}

func rollback() {
	fmt.Println("Rollback or compensation logic executed.")
}
