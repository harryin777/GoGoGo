package tests

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var val int32
	oldVal := atomic.SwapInt32(&val, 20)
	fmt.Printf("old val: %v, new val : %v \n", oldVal, val)
	atomic.StoreInt32(&val, 15)
	fmt.Println(val)

	atomic.AddInt32(&val, 10)
	fmt.Println(val)

	atomic.CompareAndSwapInt32(&val, 25, 50)
	fmt.Println(val)
}
