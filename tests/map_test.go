package tests

import (
	"log"
	"runtime"
	"testing"
)

var intMap map[int]int
var cnt = 8192

func TestClearMap(t *testing.T) {
	printMemStats()
	initMap()
	runtime.GC()
	printMemStats()
	log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	log.Println(len(intMap))
	runtime.GC()
	printMemStats()
	intMap = nil
	runtime.GC()
	printMemStats()
}
func initMap() {
	intMap = make(map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
