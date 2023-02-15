package tests

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	arrayA := [2]int{100, 200}
	testArrayPoint1(&arrayA) // 1.传数组指针
	// 注意这里, arrayB 本身是有一个新的地址没问题,但是 slice 底层是数组,slice 有三个属性,一个指针,一个长度,一个容量
	// 指针指向数组开头,也就是说这里,arrayB 指向的地址和 arrayA 指向的地址是同一个数组
	// 那么什么时候 arrayB 指向一个全新的数组, 在 arrayB 追加一个元素, 无论是在数组开始还是末尾, 但是对数组内元素的更改是不会导致指向一个新的数组,也就是会对原先的 slice 产生影响
	// 如果原先的数组还够用,那么array 追加新的元素是不会指向新的数组,只是覆盖原来数组对应的索引
	arrayB := arrayA[:]
	// 扩容, 那么最后的结果 arrayA 的第一个元素就变成了 300
	arrayB = append(arrayB, 1)
	// 所以为什么这里虽然传递的是 arrayB 的地址,但是依然会改变 arrayA 的值,因为指向的是同一个地址
	testArrayPoint2(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArrayPoint1(x *[2]int) {
	fmt.Printf("func test1 Array : %p , %v\n", x, *x)
	(*x)[0] += 100
}

func testArrayPoint2(x *[]int) {
	fmt.Printf("func test2 Array : %p , %v\n", x, *x)
	(*x)[0] += 100
}

func TestInitParameter(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	//slice2 := slice1[0:2:2]
	//fmt.Println(slice1)
	//slice2[0] = 0
	// 这时候追加会指向一个新的数组.
	//slice2 = append(slice2, 6)
	//slice2[1] = 9
	//fmt.Println(slice1)

	slice3 := slice1[0:2:3]
	fmt.Println(slice1)
	slice3[0] = 0
	// 这时候追加会不会指向一个新的数组,因为 slice3创建的时候, max 和 cap 不一致, max为切片保留的原切片的最大下标.
	slice3 = append(slice3, 6)
	// 改变了初始 slice 的值.
	slice3[1] = 9
	fmt.Println(slice1)

}

func TestQuation(t *testing.T) {
	slice := make([]int, 0, 10)
	updateSlice(slice)
	fmt.Println(slice)
}

func updateSlice(s []int) {
	s = append(s, 1)
}
