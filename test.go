package main

//标准输出包
import (
	"fmt"
)

type Person struct {
	name string
	age int
	male bool
}

func main() {
	fmt.Println("hello world")

	//定义一个多行字符串，必须用`
//	s1 := `这是一个多行
//字符串`
//
//	fmt.Println(s1)
//	fmt.Println(strings.Contains(s1, "这"))
//	//长度
//	fmt.Println(len(s1))
//	//前缀判断，HasSuffix后缀判断
//	fmt.Println(strings.HasPrefix(s1, "这是1"))
//	//字串出现的位置，以及LastIndex
//	fmt.Println(strings.Index(s1, "字"))

	//traverseStr()
	arrParameter := [4]string{"a", "q"}
	initArray(arrParameter)
	//slicePractice()

}

func traverseStr()  {
	s := "abcdefg"
	/*
		默认循环字符串结果是rune
		需要强转成string like this string(s[i])
	*/
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%v(%c) ", string(s[i]), s[i])
	}

	for _, r:= range s {
		fmt.Println(string(r))
	}

	//强制转化成数组，打印出来也是runa类型，还是需要强转
	byteS1 := []byte(s)
	for i := 0; i < len(byteS1); i++ {
		fmt.Println(string(byteS1[i]))
	}
}

func initArray(arr [4]string)  {
	//接受参数，注意如果是字符串数组，循环出来直接就是字符串，而不是数字
	arrP := arr
	for i := 0; i < len(arrP); i++ {
		fmt.Printf(arrP[i])
	}
	fmt.Println()

	//这个地方使用索引号即下标 ：指定该下标处的元素值
	arr1 := [...]int {1: 5, 6: 1}
	arr2 := arr1
	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i])
	}
	fmt.Println()
	arr1[1] = 999
	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i])
	}
	//说明新建一个数组以后内存地址是不同的
	fmt.Println(&arr1 == &arr2)

	//对象数组，注意最后一个元素是有逗号的
	target := [...]struct{
		name string
		age uint8
	}{
		{"user1", 1},
		{"user2", 10},
	}
	for i := 0; i < len(target); i++ {
		fmt.Print(target[i])
	}

	target1 := [...]Person{
		{"person1", 12, false},
	}
	for i := 0; i < len(target1); i++ {
		fmt.Print(target1[i])
	}
}

func slicePractice()  {
	array1 := [...]int{1, 2, 3, 4}
	slice1 := array1[:]
	fmt.Println(cap(slice1))
}
