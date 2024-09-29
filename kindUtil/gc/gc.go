package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var memStats runtime.MemStats
	runtime.GC() // 手动触发 GC

	runtime.ReadMemStats(&memStats)
	fmt.Printf("GC Pause Total: %v\n", memStats.PauseTotalNs)
	fmt.Printf("Last GC Pause: %v\n", memStats.PauseNs[(memStats.NumGC+255)%256])
	fmt.Printf("GC Num: %d\n", memStats.NumGC)

	// 模拟一段内存分配
	allocateMemory()

	// 等待一段时间让 GC 运行
	time.Sleep(5 * time.Second)
	runtime.GC()

	runtime.ReadMemStats(&memStats)
	fmt.Printf("GC Pause Total: %v\n", memStats.PauseTotalNs)
	fmt.Printf("Last GC Pause: %v\n", memStats.PauseNs[(memStats.NumGC+255)%256])
	fmt.Printf("GC Num: %d\n", memStats.NumGC)
}

func allocateMemory() {
	_ = make([]byte, 100<<20) // 分配 10 MB 内存
}
