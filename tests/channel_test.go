package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestNoCache(t *testing.T) {
	chan1 := make(chan int)
	fmt.Println("here 1")
	// 无缓存的 channel 会在这里阻塞住
	chan1 <- 1
	fmt.Println(len(chan1))
}

func TestCache(t *testing.T) {
	chan1 := make(chan int, 1)
	//var chan1 chan int
	fmt.Println("here 1")
	// 有缓存的 channel 就顺利可以接受一个值
	chan1 <- 1
	fmt.Println("here 2")
	// 超缓存的部分,还是会阻塞住
	chan1 <- 2
	fmt.Println(len(chan1))
}

func TestChannelNoCache(t *testing.T) {
	chan1 := make(chan int)
	_ = chan1
	// 先有监听才能在 chan 里放元素, 以下代码可以正常运行
	go func(c chan int) {
		select {
		case val := <-c:
			fmt.Println(val)
		default:
			fmt.Println("empty")
		}
	}(chan1)
	time.Sleep(1 * time.Second)
	// 会在这里阻塞,因为没有获取的地方
	chan1 <- 1
	//go func(c chan int) {
	//	select {
	//	case val := <-c:
	//		fmt.Println(val)
	//	default:
	//		fmt.Println("empty")
	//	}
	//}(chan1)
	//time.Sleep(1 * time.Second)
}

func TestChannelCache(t *testing.T) {
	chan1 := make(chan int, 1)
	_ = chan1
	// 先有监听才能在 chan 里放元素, 被注释的代码可以正常运行
	//go func(c chan int) {
	//	select {
	//	case val := <-c:
	//		fmt.Println(val)
	//	default:
	//		fmt.Println("empty")
	//	}
	//}(chan1)
	// 在这里休眠就会输出 empty, 因为 channel 是空
	//time.Sleep(1 * time.Second)
	chan1 <- 1
	go func(c chan int) {
		select {
		case val := <-c:
			fmt.Println(val)
		default:
			fmt.Println("empty")
		}
	}(chan1)
	time.Sleep(1 * time.Second)

}

func producer(c chan int, str string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("this is %v producer \n", str)
		c <- i
	}
	close(c)
}

func consumer(c chan int, str string) {
	for {
		select {
		case msg, ok := <-c:
			if ok {
				fmt.Printf("this is %v consume : %v \n", str, msg)
			} else {
				return
			}
		}
	}
}

func Test_ProConsumer(t *testing.T) {
	c := make(chan int)
	go func() {
		consumer(c, "222")
	}()
	producer(c, "111")

}
