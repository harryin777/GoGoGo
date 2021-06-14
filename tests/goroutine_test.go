/**
*   @Author: yky
*   @File: goroutine_test
*   @Version: 1.0
*   @Date: 2021-06-08 20:45
 */
package tests

import (
	"fmt"
	"math/rand"
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

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func Test_GoroutinePool(t *testing.T) {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(64, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	for i := 0; i < 2; i++ {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

/**
* @Description: 不理解用 * & 比直接传值的好处在哪里，管道也是引用类型，在函数体内的改变会影响这个值本身 TODO
* @Param:
* @return:
**/
// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
