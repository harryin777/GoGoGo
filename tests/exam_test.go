/**
*   @Author: yky
*   @File: exam_test
*   @Version: 1.0
*   @Date: 2021-06-07 21:15
 */
package tests

import (
	"fmt"
	"testing"
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
* @Description: defer 和 return的问题
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
