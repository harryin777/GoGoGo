package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex1(t *testing.T) {
	wg := sync.WaitGroup{}
	var mutex sync.Mutex
	fmt.Println("Locking  (G0)")
	mutex.Lock()
	fmt.Println("locked (G0)")
	wg.Add(3)
	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			mutex.Lock()
			fmt.Printf("locked (G%d)\n", i)
			time.Sleep(time.Second * 2)
			mutex.Unlock()
			fmt.Printf("unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	mutex.Unlock()
	fmt.Println("unlocked (G0)")
	wg.Wait()
}

const (
	// 1 左移 0 位
	mutexLocked = 1 << iota // mutex is locked
	// 1 左移 1 位
	mutexWoken
	// 1 左移 2 位
	mutexStarving
	mutexWaiterShift      = iota
	starvationThresholdNs = 1e6
)

func Test1(t *testing.T) {
	fmt.Println(mutexLocked)
	fmt.Println(mutexWoken)
	fmt.Println(mutexStarving)
	fmt.Println(mutexWaiterShift)
	fmt.Println(starvationThresholdNs)
	var a int
	a = -3
	fmt.Println(a >> 3)

}
