/**
*   @Author: yky
*   @File: goroutine_test
*   @Version: 1.0
*   @Date: 2021-06-08 20:45
 */
package tests

import (
	"fmt"
	"runtime"
	"testing"
)

/**
* @Description: 测试Gosched 是让出当前cpu时间片没错，但是为什么主协程已经执行完毕
子协程还可以继续执行，难道不应该主协程执行完之后，程序就结束了吗
* @Param:
* @return:
**/
func Test_Gosched(t *testing.T) {
	go func() {
		for i := 0; i < 40; i++ {
			fmt.Println(i)
		}
	}()
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("main", i)
	}
}

/**
* @Description: defer和 return，defer在 return后执行，defer在runtime.Goexit()之前执行
* @Param:
* @return:
**/
func Test_Goexit(t *testing.T) {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			//runtime.Goexit()
			defer fmt.Println("C.defer")
			//runtime.Goexit()
			fmt.Println("B")
			runtime.Goexit()

		}()
		fmt.Println("A")
	}()
	for {
	}
}

/**
* @Description: 无缓冲的 channel 需要先接受以后，才可以传入值
* @Param:
* @return:
**/
func Test_Channel(t *testing.T) {
	//没有指定容量的就是无缓冲的channel
	ch := make(chan int)
	go Ch2(ch)
	ch <- 10
}

func Ch2(c chan int) {
	i := <-c
	fmt.Printf("channel 获取的值 ：%v", i)
}
