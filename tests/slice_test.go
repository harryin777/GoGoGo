package tests

import (
	"fmt"
	"reflect"
	"test1/Utils"
	"testing"
	"unsafe"
)

// 关于在递归调用中， arr为什么在下层的修改对上层不可见，具体表现在arr在level = 1的时候增加了1元素，但是在level= 0的时候进行第二次递归，1元素数量还是1个
// 在递归调用中进行 append 操作时，虽然 arr 的底层数组没有变化，但 arr 切片的长度和容量在每次递归调用中是不同的。
// 因此，调用者的 arr 切片在第二次递归调用时并未包含第一个递归调用中的 append 操作的结果。
// 根据arr的地址，len，cap可以新建一个slice，这个slice里面确实有两个1元素，证明不同层的调用虽然共享了底层数组但是并不共享cap和len
func d1(arr []int, level int) {
	if len(arr) == 2 {
		return
	}
	if level == 0 {
		fmt.Printf("before first d1 : %v \n", GetSliceArrayContent(arr))
	}
	arr = append(arr, 1)
	fmt.Printf("level : %v, addr : %p \n", level, arr)
	d1(arr, level+1)
	if level == 0 {
		fmt.Printf("after first d1 : %v \n", GetSliceArrayContent(arr))
	}
	d1(arr, level+1)
}

func GetSliceArrayContent(slice interface{}) interface{} {
	sliceVal := reflect.ValueOf(slice)
	if sliceVal.Kind() != reflect.Slice {
		return nil
	}

	// Get the address of the slice's underlying array
	arrayPtr := unsafe.Pointer(sliceVal.Pointer())

	// Create a slice header for the underlying array
	arrayHeader := reflect.SliceHeader{
		Data: uintptr(arrayPtr),
		Len:  sliceVal.Cap(),
		Cap:  sliceVal.Cap(),
	}

	// Create a slice from the array header
	arrayVal := reflect.NewAt(reflect.SliceOf(sliceVal.Type().Elem()), unsafe.Pointer(&arrayHeader)).Elem()
	return arrayVal.Interface()
}

func Test_d1(t *testing.T) {
	arr := make([]int, 0, 10)
	d1(arr, 0)
}

func TestPointer(t *testing.T) {
	// 这是个数组
	arrayA := [3]int{100, 200, 300}
	testArrayPoint1(&arrayA) // 1.传数组指针
	// 注意这里, arrayB 本身是有一个新的地址没问题,但是 slice 底层是数组,slice 有三个属性,一个指针,一个长度,一个容量
	// 指针指向数组开头,也就是说这里,arrayB 指向的地址和 arrayA 指向的地址是同一个数组
	// 那么什么时候 arrayB 指向一个全新的数组, 在 arrayB 追加一个元素, 无论是在数组开始还是末尾, 但是对数组内元素的更改是不会导致指向一个新的数组,也就是会对原先的 slice 产生影响
	// 如果原先的数组还够用,那么array 追加新的元素是不会指向新的数组,只是覆盖原来数组对应的索引
	arrayB := arrayA[:1:2]
	fmt.Printf("截取之后 : %v \n", arrayB)
	// 扩容,而且扩容必须要超过B的capacity，B指向了新的地址，所以最后A的0位置元素不是300，因为B已经和A指向的不是同一个地址了，如果注释掉这一段，还是300
	arrayB = append(arrayB, 1, 1)
	// 不会扩容
	//arrayB = append(arrayB, 1)
	// 所以为什么这里虽然传递的是 arrayB 的地址,但是依然会改变 arrayA 的值,因为指向的是同一个地址
	testArrayPoint2(&arrayB) // 2.传切片
	testArrayPoint3(arrayB)  // 2.传切片
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

func testArrayPoint3(x []int) {
	fmt.Printf("func test3 Array : %p , %v\n", x, x)
	x[0] += 100
}

/*
change 函数的参数s是原始切片slice的一个副本，它们共享同一个底层数组，但是切片的长度和容量信息是各自独立的
第一次进入 change, 参数 s 的长度和容量分别是 3,5, append 操作不会引发扩容, 所以不会指向新的数组, 底层数组从原来的 1,2,0,0,0 变成了 1,2,0,3,4
第二次进入 change, 参数 s 的长度和容量分别是 2,2, append 操作会引发扩容, 所以会指向新的数组, 底层数组从原来的 1,2,0,3,4 变成了 1,2,3,4,0
外层的 slice 指向新的数组, 但是长度和容量和change 内的 s 是独立的
*/
func TestAppend(t *testing.T) {
	slice := make([]int, 3, 5)
	slice[0] = 1
	slice[1] = 2
	fmt.Printf("slice 1: %p , %v\n", slice, slice)
	//slice = append(slice, 3)
	change(slice)
	fmt.Printf("slice 2: %p , %v\n", slice, slice)
	change(slice[0:2])
	fmt.Printf("slice 3: %p , %v\n", slice, slice)

}

func change(s []int) {
	fmt.Printf("change before: %p , %v\n", s, s)
	s = append(s, 3)
	s = append(s, 4)
	fmt.Printf("change after: %p , %v\n", s, s)
}

func TestInitParameter(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	// 注意这里已经指定了容量为 2
	slice2 := slice1[0:2:2]
	fmt.Printf("1 slice1:%v \n", slice1)
	fmt.Printf("2 slice2:%v \n", slice2)
	// 这时候追加会指向一个新的数组. 因为原来的容量是2，追加需要扩容，所以slice2就指向新的数组地址
	slice2 = append(slice2, 6)
	slice2[0] = 0
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
	c := []int{1, 2}
	_ = c
	fmt.Println(len(a))
	fmt.Println(a == b)
	// c 是一个 slice a,b 都是数组
	//fmt.Println(a == c)
}

/*
	模拟了一个一对多的场景,A 对 B,在乱序的情况下,map 里存放了 A 的地址,后续遍历的时候,如果 A 的地址相同,则认为是同一个 A

在这个地址追加 B,看是否可以成功,是可以的.注意的点是for 循环 val 的使用,记住 val 是一个临时地址
*/
func Test_addressSlice(t *testing.T) {
	type B struct {
		Name string
	}
	type A struct {
		Id int
		B  []B
	}
	s1 := make([]*A, 0, 10)
	s2 := []A{
		{
			Id: 1,
			B: []B{
				{Name: "1"},
			},
		},
		{
			Id: 2,
			B: []B{
				{Name: "2"},
			},
		},
		{
			Id: 1,
			B: []B{
				{Name: "1"},
			},
		},
	}
	AMapInfo := make(map[int]*A)
	for _, val := range s2 {
		if info, ok := AMapInfo[val.Id]; ok {
			info.B = append(info.B, val.B...)
		} else {
			tmp := &A{
				Id: val.Id,
				B:  val.B,
			}
			s1 = append(s1, tmp)
			AMapInfo[val.Id] = tmp
		}
	}

	Utils.ReceiveStruct(s1)
}

type ex struct {
	Id int
}

func Test_forRange(t *testing.T) {
	var eslice []*ex
	eslice = append(eslice, &ex{Id: 1}, &ex{Id: 2}, &ex{Id: 3})
	for _, e := range eslice {
		e.Id = 100
	}
	for _, e := range eslice {
		fmt.Println(e.Id)
	}
}
