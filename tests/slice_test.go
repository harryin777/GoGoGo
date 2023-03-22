package tests

import (
	"fmt"
	"sort"
	"testing"
)

func TestPointer(t *testing.T) {
	arrayA := [3]int{100, 200, 300}
	testArrayPoint1(&arrayA) // 1.传数组指针
	// 注意这里, arrayB 本身是有一个新的地址没问题,但是 slice 底层是数组,slice 有三个属性,一个指针,一个长度,一个容量
	// 指针指向数组开头,也就是说这里,arrayB 指向的地址和 arrayA 指向的地址是同一个数组
	// 那么什么时候 arrayB 指向一个全新的数组, 在 arrayB 追加一个元素, 无论是在数组开始还是末尾, 但是对数组内元素的更改是不会导致指向一个新的数组,也就是会对原先的 slice 产生影响
	// 如果原先的数组还够用,那么array 追加新的元素是不会指向新的数组,只是覆盖原来数组对应的索引
	arrayB := arrayA[:1:2]
	// 扩容,而且扩容必须要超过B的capacity，B指向了新的地址，所以最后A的0位置元素不是300，因为B已经和A指向的不是同一个地址了，如果注释掉这一段，还是300
	arrayB = append(arrayB, 1, 1)
	// 所以为什么这里虽然传递的是 arrayB 的地址,但是依然会改变 arrayA 的值,因为指向的是同一个地址
	testArrayPoint2(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)
}

func testArrayPoint1(x *[3]int) {
	fmt.Printf("func test1 Array : %p , %v\n", x, *x)
	(*x)[0] += 100
}

func testArrayPoint2(x *[]int) {
	fmt.Printf("func test2 Array : %p , %v\n", x, *x)
	(*x)[0] += 100
}

func TestInitParameter(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[0:2:2]
	fmt.Printf("1 slice1:%v \n", slice1)
	fmt.Printf("2 slice2:%v \n", slice2)
	slice2[0] = 0
	// 这时候追加会指向一个新的数组. 因为原来的容量是2，追加需要扩容，所以slice2就指向新的数组地址
	slice2 = append(slice2, 6)
	fmt.Printf("3 slice1:%v \n", slice1)
	fmt.Printf("4 slice2:%v \n", slice2)
	slice2[1] = 9
	fmt.Printf("5 slice1:%v \n", slice1)
	fmt.Printf("6 slice2:%v \n", slice2)

	slice3 := slice1[0:2:3]
	fmt.Printf("7 slice3:%v \n", slice3)
	slice3[0] = 66
	// 这时候追加会不会指向一个新的数组,因为 slice3创建的时候, max 和 cap 不一致, max为切片保留的原切片的最大下标.
	// 也就是容量是3，还可以再追加一个，只不过在原数组基础上修改
	slice3 = append(slice3, 6)
	fmt.Printf("8 slice1:%v \n", slice1)
	// 改变了初始 slice 的值.
	slice3[1] = 9
	fmt.Printf("9 slice1:%v \n", slice1)
	fmt.Printf("10 slice3:%v \n", slice3)

}

func TestQuation(t *testing.T) {
	slice := make([]int, 0, 10)
	updateSlice(slice)
	fmt.Println(slice)
}

func updateSlice(s []int) {
	s = append(s, 1)
}

func TestSortSlice(t *testing.T) {
	slice1 := []int{1, 5, 6, 20, 4, 2}
	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})

	fmt.Println(slice1)
}

func TestShallowCopy(t *testing.T) {
	a := []int{1, 2, 3}
	//b := a
	// 是浅拷贝
	//b[0] = 0
	//fmt.Println(a)
	changeSlice(a)
	fmt.Println(a)
}

// 引用传递
func changeSlice(s []int) {
	s[0] = 0
}

func TestZuHe(t *testing.T) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			fmt.Printf("%v,%v \n", i, j)
		}
	}
}

// 二维slice的正确深拷贝！！！！！！！！
func TestDeepCopy(t *testing.T) {
	tmp1 := [][]int{{1}, {2}, {3}}
	tmp2 := make([][]int, 3, 3)
	for i := 0; i < len(tmp2); i++ {
		tmp2[i] = make([]int, 1, 1)
		copy(tmp2[i], tmp1[i])
	}

	tmp2[0][0] = 0
	fmt.Println(tmp2)
	fmt.Println(tmp1)
}

// 指定索引初始化slice
func TestAssignIndex(t *testing.T) {
	// 指定了索引，但是前后两个位置冲突了，所以报错
	//a := []int{2: 2, 3, 1: 5, 6}
}

func TestInitArray(t *testing.T) {
	a := [...]int{1, 2}
	b := [2]int{1, 2}
	fmt.Println(len(a))
	fmt.Println(a == b)
}
