/**
*   @Author: yky
*   @File: exam_test
*   @Version: 1.0
*   @Date: 2021-06-07 21:15
 */
package tests

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"test1/Utils"
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
	count := 0
	var lock sync.Mutex
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		thread1 := "thread1"
		for i := 0; count < 99; i++ {
			lock.Lock()
			count = count + 1
			fmt.Printf("goroutine : %v, count : %v \n", thread1, count)
			time.Sleep(1)
			lock.Unlock()
		}
		wg.Done()
	}()

	go func() {
		thread1 := "thread2"
		for i := 0; count < 99; i++ {
			lock.Lock()
			count = count + 1
			fmt.Printf("goroutine : %v, count : %v \n", thread1, count)
			time.Sleep(1)
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
}

func Test_QueryHtml(t *testing.T) {
	//https://www.iapco.org/publications/on-line-dictionary/dictionary/?ds=visitor

	words := []string{
		"exhibitor",
		"main exhibitor ",
		"co-exhibitor ",
		"represented company",
		"international exhibitor",
		"foreign exhibitor",
		"national exhibitor",
		"domestic exhibitor",
		"exhibitor staff",
		"exhibitor personnel",
		"visitor",
		"trade visitor",
		"general public visitor",
		"international visitor",
		"foreign visitor",
		"national visitor",
		"domestic visitor",
		"visit",
		"hosted visitor",
		"delegate",
		"international delegate",
		"foreign delegate",
		"national delegate",
		"domestic delegate",
		"accompanying person",
		"media representative",
		"service provider",
		"official contractor",
		"sponsor",
		"organizer",
		"co-organizer",
		"show management",
		"attendee",
		"total attendance",
		"admission category",
		"Types of events",
		"exhibition",
		"show",
		"fair",
		"trade exhibition",
		"international exhibition",
		"public exhibition",
		"general exhibition",
		"specialized exhibition",
		"conference",
		"convention",
		"seminar",
		"symposium",
		"workshop",
		"Physical item",
		"booth stand",
		"booth space/stand space",
		"raw space",
		"contra booth",
		"contra stand",
		"pavilion",
		"gross indoor exhibition venue space",
		"gross outdoor exhibition venue space",
		"gross exhibition space",
		"net exhibition space",
		"rented exhibition space",
		"floor plan",
		"exhibitors' manual",
		"exhibition directory",
		"exhibition catalogue",
		"convention centre",
		"congress centre",
		"exhibition centre",
		"fairground",
		"Miscellaneous",
		"build up",
		"tear down",
		"break down",
		"duration of exhibition",
	}

	f := Utils.CreateFile("./fileThing.txt")
	for _, w := range words {
		time.Sleep(1000)
		request, err := http.NewRequest("GET", fmt.Sprintf("https://www.iapco.org/publications/on-line-dictionary/dictionary/?ds=%v", w), nil)
		if err != nil {
			log.Fatal(err)
		}
		client := &http.Client{}

		resp, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		var body []byte
		_ = body
		if resp.Header.Get("Content-Encoding") == "gzip" {
			println("--------------------gzip")
			res := new(bytes.Buffer)
			gr, err := gzip.NewReader(resp.Body)
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(res, gr)
			if err != nil {
				panic(err)
			}
			body = res.Bytes()
		} else {
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
		}

		strHtml := string(body)

		//这种方式也可以获取到
		//goquery.NewDocumentFromReader(resp.Body)
		gq, err := goquery.NewDocumentFromReader(strings.NewReader(strHtml))
		if err != nil {
			panic(err)
		}

		buf := bufio.NewWriter(f)
		// .是 class 筛选,#是 id 筛选
		gq.Find("body").Find(".search-result").Each(func(i int, selection *goquery.Selection) {
			word := selection.Find(".search-result__word").Text()
			des := selection.Find(".search-result__description").Text()
			_, err := buf.WriteString(fmt.Sprintf("word: %v \n", word))
			if err != nil {
				panic(err)
			}
			_, err = buf.WriteString(fmt.Sprintf("des: %v ", des))
			if err != nil {
				panic(err)
			}
		})
	}
}

/*
以下代码输出什么 % 运算只能用于 整数类型。1 % 2.0，两个操作数都是字面量常量，都是无类型的
这时会以 2.0 的 untype float constant 为准，1 隐式转为 untype float constant，所以编译错误。
而 int(1) % 2.0 中，2.0 是无类型的，int(1) 是 int，因此 2.0 会转为 int，因此能正常编译。
*/
//func Test_Exam8(t *testing.T) {
//	fmt.Println(1 % 2.0)
//	fmt.Println(int(1) % 2.0)
//}
