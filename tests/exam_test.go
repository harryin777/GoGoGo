/**
*   @Author: yky
*   @File: exam_test
*   @Version: 1.0
*   @Date: 2021-06-07 21:15
 */
package tests

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
	"time"
)

/**
* @Description: Exam1 会输出什么
* @Param:
* @return:
**/
func Test_Exam1(t *testing.T) {
	a := make([]int, 20)
	a = []int{7, 8, 9, 10}
	b := a[15:16]
	fmt.Println(b)
}

/**
* @Description: 循环内使用新值
* @Param:
* @return:
**/
func Test_Exam2(t *testing.T) {
	in := []int{1, 2, 3}

	var out []*int
	for _, v := range in {
		//v := v
		out = append(out, &v)
	}
	//输出结果最后是一样的，因为 v 作为一个变量，他的地址一直不变，out中存入的是 v 的地址
	//这个地址存储的内容每次都是不同的，直到最后一次是3，然后存进去的三个地址都指向 3
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}

/**
* @Description: defer 和 return的问题，命名返回值和非命名返回值
* @Param:
* @return:
**/
func Test_Exam3(t *testing.T) {
	i := 10
	//这里传入的是地址，所以函数体内对变量的改变会修改变量的值
	j := hello(&i)
	fmt.Println(i, j)
}

func hello(i *int) int {
	defer func() {
		*i = 19
	}()
	return *i
}

/**
* @Description: 主协程退出，子协程就会停止执行
* @Param:
* @return:
**/
func Test_Exam4(t *testing.T) {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

type exam51 struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

type exam52 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Test_Exam5(t *testing.T) {

	e1 := exam51{
		Name: "e5",
		Age:  0,
	}
	eb1, _ := jsoniter.Marshal(e1)
	fmt.Printf("e1:%s \n", eb1)

	e2 := exam52{
		Name: "e5",
	}
	eb2, _ := jsoniter.Marshal(e2)
	fmt.Printf("e2:%s \n", eb2)
}
