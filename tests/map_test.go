package tests

import (
	"fmt"
	"log"
	"runtime"
	"test1/Utils"
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

type student struct {
	Name string
	Age  int
}

func TestRange(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		// 可以看到地址一直没有变化,也就是 stu 这个变量只初始化一次,后续都是在前面的同一个地址上赋值,所以 val 不变,但是地址一直都是一样.
		// 那么最后 map 里的值也都是一样的,因为是从同一个内存地址获取的.
		//fmt.Printf("%p \n", &stu)
		// 为什么这种操作可以,这是新建了一个临时局部变量,会给这个临时变量分配新的地址,值是 stu 的值,那么 map 最后指向的是临时变量的地址对应的值
		// tmp := stu
		m[stu.Name] = &stu
	}
	//map 中存放的是地址
	//for _, val := range m {
	//	val.Age = 999
	//}
	Utils.ReceiveStruct(m)

	m2 := make(map[string]student)
	for _, stu := range stus {
		m2[stu.Name] = stu
	}
	//map 中存放的是值
	for _, val := range m2 {
		val.Age = 999
	}
	fmt.Printf("map2 : %v \n", m2)

}

func TestMapAndSlice(t *testing.T) {
	// 测试下 slice, 初始化容量,防止扩容
	arr1 := []int{1, 2, 3, 4, 5}
	slice1 := make([]int, 0, 0)
	map1 := make(map[int]int)

	for _, val := range arr1 {
		slice1 = append(slice1, val)
		map1[val] = val
	}

	fmt.Printf("slice2 : %v \n", slice1)
	fmt.Printf("map2 : %v \n", map1)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	slice2 := make([]student, 0, 0)
	map2 := make(map[string]student)

	for _, val := range stus {
		slice2 = append(slice2, val)
		map2[val.Name] = val
	}

	fmt.Printf("slice2 : %v \n", slice2)
	fmt.Printf("map2 : %v \n", map2)
}

var minVal = ^(int(^uint(0) >> 1))
var maxVal = int(^uint(0) >> 1)

func TestInitCapMap(t *testing.T) {
	fmt.Println(maxVal)
	fmt.Println(minVal)
	mm := make(map[int]int, 10)
	for key, val := range mm {
		fmt.Println(key)
		fmt.Println(val)
	}
}
