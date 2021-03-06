/**
*   @Author: yky
*   @File: timer_test
*   @Version: 1.0
*   @Date: 2021-06-16 21:52
 */
package tests

import (
	"fmt"
	"testing"
	"time"
)

/**
* @Description: 只执行一次的定时器，而且定时器的执行是同步阻塞的，需要避免阻塞，除非本来就需要阻塞
* @Param:
* @return:
**/
func Test_OnceTimer(t *testing.T) {
	go timer1()
	go timer2()
	fmt.Println("will end")
	time.Sleep(5 * time.Second)
}

func testTimer1() {
	go func() {
		fmt.Println("test timer1")
	}()

}

func testTimer2() {
	go func() {
		fmt.Println("test timer2")
	}()
}

//多次循环定时器 ticker
func timer1() {
	timer1 := time.NewTicker(1 * time.Second)
	//select {
	//case <-timer1.C:
	//	testTimer1()
	//}
	for {
		<-timer1.C
		testTimer1()
	}

}

//单次循环定时器
func timer2() {
	timer2 := time.NewTimer(3 * time.Second)

	select {
	case <-timer2.C:
		testTimer2()
	}

}

func Test_DateTimestamp(t *testing.T) {
	var ts int64
	ts = 1631271819000
	fmt.Println(time.Unix(ts/1000, 0))
}
