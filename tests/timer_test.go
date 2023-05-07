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

//buf := bufio.NewReader(os.Stdin)
//var data string
//for {
//str, err := buf.ReadString('\n')
//if err != nil || err == io.EOF {
//break
//}
//data  = strings.Trim(str, "\n")
//}
//arr := strings.Split(data, " ")
//arr = arr[:len(arr)-1]

func TestN1(t *testing.T) {
	arr := []string{"268", "90", "179", "129", "204", "224"}
	_ = arr
	list := []int{268, 90, 179, 129, 204, 224}
	dp := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[i] > list[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	var dataa int
	for i := 0; i < len(dp); i++ {
		dataa = max(dataa, dp[i])
	}

	fmt.Println(dataa)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func Test102Str(t *testing.T) {
	sum := 0
	str := "97"
	for _, val := range str {
		sum = sum*10 + (int(val - '0'))
	}
	fmt.Println(sum <= 255)
}

func TestGetLastMonth(t *testing.T) {
	lastMonth := time.Now().AddDate(0, -1, 0)
	t2 := lastMonth.Format("200601")
	month := fmt.Sprintf("%d%d", lastMonth.Year(), int(lastMonth.Month()))
	fmt.Println(month)
	fmt.Println(t2)
}
