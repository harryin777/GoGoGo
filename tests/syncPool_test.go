package tests

import (
	"fmt"
	"sync"
	"testing"
)

var pool *sync.Pool

func InitPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			return 1
		},
	}
}

func InitPool2() {
	pool = &sync.Pool{
		New: func() interface{} {

			return Person{
				Name: "test",
			}
		},
	}
}

func InitPool3() {
	pool = &sync.Pool{
		New: func() interface{} {

			return &Person{
				Name: "test",
			}
		},
	}
}

func TestGetDirectly(t *testing.T) {
	InitPool()
	var wg sync.WaitGroup
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(pool.Get())
		}(&wg)
	}
	wg.Wait()
}

func TestGetDirectlyStruct(t *testing.T) {
	InitPool2()
	var wg sync.WaitGroup
	count := 3
	chan1 := make(chan int)
	wg.Add(count)
	//pool.Put(Person{
	//	Name: "111",
	//	Age:  111,
	//	Male: false,
	//})
	for i := 0; i < count; i++ {
		t := i
		go func(wg *sync.WaitGroup, count int, chan1 chan int) {
			defer wg.Done()

			var tmp Person
			if count == 1 {
				tmp.Age = 12
				pool.Put(tmp)
				fmt.Println(1)
				chan1 <- 1
			} else if count == 2 {
				<-chan1
				fmt.Println(2)
				tmp = pool.Get().(Person)
			}
			fmt.Printf("tmp : %v , count : %v \n", tmp, count)
		}(&wg, t, chan1)
	}
	wg.Wait()
}

func TestGetDirectlyStructPointer(t *testing.T) {
	InitPool3()
	var wg sync.WaitGroup
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		t := i
		go func(wg *sync.WaitGroup, count int) {
			defer wg.Done()

			tmp := pool.Get().(*Person)
			if count == 1 {
				tmp.Age = 12
				pool.Put(tmp)
				fmt.Println(1)
			}
			fmt.Printf("tmp : %v \n", tmp)
		}(&wg, t)
	}
	wg.Wait()
}

func InitPool4() {
	pool = &sync.Pool{
		New: func() interface{} {
			return new(Person)
		},
	}
}

// sync.pool 只能在当前协程里保留改变?
func TestGetDirectlyStructPointerNew(t *testing.T) {
	InitPool4()
	var wg sync.WaitGroup
	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		t := i
		go func(wg *sync.WaitGroup, count int) {
			defer wg.Done()

			tmp := pool.Get().(*Person)
			if count == 1 {
				tmp.Age = 12
				pool.Put(tmp)
			}
			fmt.Printf("tmp : %v , count : %v \n", tmp, count)
		}(&wg, t)
	}
	wg.Wait()
}
