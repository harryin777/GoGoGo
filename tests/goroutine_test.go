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
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

/*
*
* @Description: 测试Gosched 是让出当前cpu时间片没错，但是为什么主协程已经执行完毕
子协程还可以继续执行，难道不应该主协程执行完之后，程序就结束了吗
* @Param:
* @return:
*
*/
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
	fmt.Println(2)
	ch <- 10
}

func Ch2(c chan int) {
	fmt.Println(1)
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
	// 这里必须传对象地址
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
	// TODO 主程执行完了为什么协程还可以打印
	for i := 0; i < 5; i++ {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
	time.Sleep(5000)
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

var x1 int64
var wg1 sync.WaitGroup
var lock1 sync.Mutex

/**
* @Description: 互斥锁，waitGroup类比 countdownLatch
* @Param:
* @return:
**/
func Test_Lock(t *testing.T) {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	fmt.Println(x1)
}

func add() {
	for i := 0; i < 5000; i++ {
		lock1.Lock() // 加锁
		x1 = x1 + 1
		lock1.Unlock() // 解锁
	}
	wg1.Done()
}

var (
	x2      int64
	wg2     sync.WaitGroup
	lock2   sync.Mutex
	rwlock2 sync.RWMutex
)

/**
* @Description: 读写锁
* @Param:
* @return:
**/
func Test_RWLock(t *testing.T) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go read()
	}

	wg2.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock2.Lock() // 加写锁
	x2 = x2 + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock2.Unlock()                  // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg2.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock2.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock2.RUnlock()            // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg2.Done()
}

func Test_SyncMap(t *testing.T) {
	//直接声明也可以用，不需要像 map那样一定要make
	s1 := new(sync.Map)
	s1.Store("k1", "v1")
	s1.Store("k2", "v2")
	s1.Store("k3", "v3")

	//无序遍历
	s1.Range(func(key, value interface{}) bool {
		fmt.Printf("%v, %v \n", key, value)
		return true
	})
}

var (
	count int32
	wg    sync.WaitGroup
)

/*
*
* @Description: atomic 包中的内容
留意这里atomic.LoadInt32和atomic.StoreInt32两个函数，一个读取int32类型变量的值，
一个是修改int32类型变量的值，这两个都是原子性的操作
* @Param:
* @return:
*
*/
func Test_Atomic(t *testing.T) {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count, value)
	}
}

/*
*

	 @Description 有缓存的管道,有缓存的管道在最后被调用时,主线程会等待子线程,不会直接结束
	 @Param
	 @return
	*
*/
func Test_cacheChannel(t *testing.T) {
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程
	for i := 0; i < cap(done); i++ {
		go func() {
			fmt.Println("hello world")
			done <- 1
		}()
	}

	// 等待N个后台线程完成
	//for i := 0; i < cap(done); i++ {
	//	<-done
	//}
}

// sync: negative WaitGroup counter 这个报错,就是多 done 了,导致 wg.add 的已经 done 成了负数
func Test_wgDone(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 5; i++ {
		count := i
		go func() {
			defer wg.Done()
			fmt.Printf("lala this is : %v \n", count)
		}()
	}
	wg.Wait()
	fmt.Println("finish")
}

func Test_WhyIsThereSoManyGorountines(t *testing.T) {
	go func() {

		var wg sync.WaitGroup
		wg.Add(5)
		for i := 0; i < 5; i++ {
			go func() {
				fmt.Printf("num : %d", i)
			}()
		}
		wg.Wait()
	}()
}

func Test_2Goroutines(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(wg *sync.WaitGroup, ch1, ch2 chan int) {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			select {
			case <-ch1:
				fmt.Println(i)
				ch2 <- 2
			}
		}
	}(&wg, ch1, ch2)

	go func(wg *sync.WaitGroup, ch1, ch2 chan int) {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			select {
			case <-ch2:
				fmt.Printf("%c \n", i+97)
				if i != 1 {
					ch1 <- 1
				}
			}
		}
	}(&wg, ch1, ch2)
	ch1 <- 1
	wg.Wait()

}

func Test_2G2(t *testing.T) {
	ch := make(chan int)
	go printLetters(ch)
	printNumbers(ch)
}

func printLetters(c chan int) {
	for _, val := range "abcdefg" {
		<-c
		fmt.Println(string(val))
		c <- 1
	}
}

func printNumbers(c chan int) {
	for i := 0; i < 7; i++ {
		fmt.Println(i)
		c <- 1
		<-c
	}
}

func Test_2G22(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	go func(c chan int, wg *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		wg.Done()
	}(c, &wg)

	go func(c chan int, wg *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			val := <-c
			fmt.Println(val)
		}
		wg.Done()
	}(c, &wg)
	wg.Wait()
}
