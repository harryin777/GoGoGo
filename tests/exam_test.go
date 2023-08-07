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
	"math"
	"sync"
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

/**
 * @Description 只有不具名的返回值 defer 才不能改变,如果是具名的返回值,defer 会修改最后的结果
 * @Param
 * @return
 **/
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

func f(a ...int) {
	fmt.Printf("%#v\n", a)
}

/**
 * @Description 应该输出什么?
 * @Param
 * @return
 **/
func Test_Exam6(t *testing.T) {
	f()
}

/*
Test_Exam7 循环打印 1-100 用两个线程
*/
func Test_Exam7(t *testing.T) {
	c1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i += 2 {
			c1 <- 1
			fmt.Println(i)
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i < 10; i += 2 {
			<-c1
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
	select {}
}

func Test_Exam7_1(t *testing.T) {
	//count := 0
	//var lock sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(2)
	c := make(chan int)
	go func() {
		thread1 := "thread1"
		for i := 0; i < 100; i++ {
			//lock.Lock()
			c <- 1
			if i%2 == 0 {
				fmt.Printf("goroutine : %v, count : %v \n", thread1, i)
			}
			//time.Sleep(1)
			//lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		thread1 := "thread2"
		for i := 0; i < 100; i++ {
			//lock.Lock()
			<-c
			if i%2 != 0 {
				fmt.Printf("goroutine : %v, count : %v \n", thread1, i)
			}
			//time.Sleep(1)
			//lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
}

/*
以下代码输出什么 % 运算只能用于 整数类型。1 % 2.0，两个操作数都是字面量常量，都是无类型的
这时会以 2.0 的 untype float constant 为准，1 隐式转为 untype float constant，所以编译错误。
而 int(1) % 2.0 中，2.0 是无类型的，int(1) 是 int，因此 2.0 会转为 int，因此能正常编译。
*/
//func Test_Exam8(t *testing.T) {
//	fmt.Println(1 % 2.0)
//	fmt.Println(int(1) % 2.0)
//

func Test_Exam9(t *testing.T) {
	x := math.Inf(1)
	switch {
	case x < 0, x > 0:
		fmt.Println(x)
	case x == 0:
		fmt.Println("zero")
	default:
		fmt.Println("something else")
	}
	a := int(^uint(0) >> 1)
	fmt.Println(a)
	fmt.Println(float64(a) > x)
}

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	fn := func(i int) {
		fmt.Printf("here : %v \n", i)
		ch <- vs[i](name)
	}
	fmt.Println(len(vs))
	for i, _ := range vs {
		go fn(i)
	}
	return <-ch
}

func Test10(t *testing.T) {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})

	fmt.Println(ret)
}

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

// 不同的是*Student 的定义后本身没有初始化值，所以 *Student 是 nil的，但是*Student 实现了 People 接口，接口不为 nil。
func Test11(t *testing.T) {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	// 这里是编译失败的 *Student实现了People，不是Student
	//var peo People = Student{}
}

type Person1 struct {
	age int
}

// 以下代码输出啥
func Test12(t *testing.T) {
	//person := Person1{28}
	person := &Person1{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person1) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

func Test13(t *testing.T) {
	var i *int
	//i = new(int)
	*i = 10
	fmt.Println(*i)
}

type People1 interface {
	Speak(string) string
}

type Student1 struct{}

// 注意这里，是指针类型，如果想编译成功，需要去掉*
func (stu *Student1) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func TestExam14(t *testing.T) {
	// 如果是 Student1{} 编译失败，如果是&Student1{} 可以正常编译
	var peo People1 = &Student1{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func TestExam15(t *testing.T) {
	//fmt.Println(test11())
	fmt.Println(test22())
	fmt.Println(test3())
	fmt.Println(test4())

	return
}

func test11() (v int) {
	defer fmt.Println(v)
	return v
}

func test22() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}

func TestExam16(t *testing.T) {
	a := []int{1, 1, 2, 2, 3, 4, 4}
	ans := 0
	for i := 0; i < len(a); i++ {
		ans = ans ^ a[i]
	}
	fmt.Println(ans)
}

type People17 struct{}

func (p *People17) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People17) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People17
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Test17(t *testing.T) {
	t1 := Teacher{}
	t1.ShowB()
}

func Test18(t *testing.T) {
	var a uint8 = 1
	var b uint8 = 255
	fmt.Println("减法：", a-b)
	fmt.Println("加法：", a+b)
	fmt.Println("乘法：", a*b)
}

func Test19(t *testing.T) {
	a := ^uintptr(0)
	b := a >> 63
	c := 4 << b
	d := c * 8
	e := d - 1
	fmt.Printf("a:%d, b:%d, c:%d, d:%d, e:%d", a, b, c, d, e)
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	fmt.Printf("t address : %p \n", &t)
	return c
}

func Test20(t *testing.T) {
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
	fmt.Println(a("All"))
}

// 注意删除了 A 并且第一次循环map 里获取的不是 A 那么 counter 会少加一次
func Test21(t *testing.T) {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

func Test22(t *testing.T) {
	total, sum := 0, 0
	for i := 1; i <= 10; i++ {
		sum += i
		go func() {
			total += i
		}()
	}
	fmt.Printf("total:%d sum %d", total, sum)
}

func f23(n int) (r int) {
	defer func() {
		r += n
		//recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func Test23(t *testing.T) {
	fmt.Println(f23(3))
}
