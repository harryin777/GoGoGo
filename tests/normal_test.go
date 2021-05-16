package tests

//标准输出包
import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
	Male bool
}

func Test_TraverseStr(t *testing.T) {
	s := "abcdefg"
	/*
		默认循环字符串结果是rune
		需要强转成string like this string(s[i])
	*/
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		fmt.Printf("%v(%c) ", string(s[i]), s[i])
	}
	// _代表一个占位符，range循环字符串的时候，前面两个参数是position，character
	for _, r := range s {
		fmt.Printf("%+v", r)
	}

	//强制转化成数组，打印出来也是runa类型，还是需要强转
	byteS1 := []byte(s)
	for i := 0; i < len(byteS1); i++ {
		fmt.Println(string(byteS1[i]))
	}
}

func Test_InitArray(t *testing.T) {
	//接受参数，注意如果是字符串数组，循环出来直接就是字符串，而不是数字
	arrP := [4]int{
		1,
		5,
	}
	for i := 0; i < len(arrP); i++ {
		//fmt.Printf(string(arrP[i]))
	}
	fmt.Println()

	//指定下标元素值，初始化并新建数组
	arr1 := [...]int{1: 5, 6: 1}
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

	//初始化结构体数组
	target := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 1},
		{"user2", 10},
	}
	//普通循环结构体数组
	for i := 0; i < len(target); i++ {
		fmt.Print(target[i])
	}
	for key, val := range target {
		fmt.Printf("key:%v, val: %v\n", key, val)
	}

}

func Test_SlicePractice(t *testing.T) {
	array1 := [...]int{1, 2, 3, 4}

	slice1 := array1[:]
	for i := 0; i < len(slice1); i++ {
		fmt.Print(slice1[i])
	}
	fmt.Println()

	slice2 := array1[:]
	slice1[1] = 1000
	for i := 0; i < len(slice1); i++ {
		fmt.Print(slice1[i])
	}
	//内存地址不同，但是值却是相同的
	fmt.Println(&slice1 == &slice2)
	for i := 0; i < len(slice2); i++ {
		fmt.Print(slice2[i])
	}

	//元素为map类型的切片，注意一定在最前面有个[]
	slice3 := make([]map[string]string, 8)
	slice3[0] = make(map[string]string, 2)
	slice3[0]["key1"] = "val1"

	//测试只声明key
	for k := range slice3 {
		fmt.Println(k)
	}

	//测试append一个slice给另一个slice
	slice4 := []int{7, 77, 777}
	slice4 = append(slice4, slice1...)

}

/**
make(map[KeyType]ValueType, [cap])
*/
func Test_Map(t *testing.T) {

	userInfo := map[string]string{
		"username": "aaa",
		"password": "123",
	}
	fmt.Println(userInfo)
	//判断键值是否存在，ok存在为true，不存在为false，v为值类型的零值
	v, ok := userInfo["username1"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("nothing at all")
	}
	//map的遍历
	for k, v := range userInfo {
		fmt.Printf("%v\n %v\n", k, v)
	}

	//value为slice的map
	mapContainSlice := make(map[string][]string, 3)
	//初始化aaa对应的value值，长度为8
	mapContainSlice["aaa"] = make([]string, 8)
	fmt.Println(len(mapContainSlice["aaa"]))
}

type tem struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Test_Struct(t *testing.T) {
	name := "aaa"
	age := 18
	//新建结构体的语法
	person := tem{Name: name, Age: age}
	fmt.Println(person)

	//匿名结构体
	var temp struct {
		Name string
		Age  int
	}
	temp.Name = "bbb"
	temp.Age = 18
	fmt.Println(temp)

	//这种初始化方式必须初始化所有的字段
	p1 := tem{
		"p1",
		12,
	}
	fmt.Printf("p7=%#v\n", p1)

	//使用键值对初始化可以初始化部分字段，其他字段保留零值
	p2 := tem{
		Name: "age",
	}
	fmt.Printf("p7=%#v\n", p2)
}

/**
这是一个方法
函数和方法的区别在于 函数是没有接受者，也就是名称前面那个参数
方法的调用和函数的不同在于
函数可以直接调用，方法必须要实例化接受者以后才可以调用
*/
func (t *tem) adj1(newAge int) {
	t.Age = newAge
}

func (t tem) adj2(newAge int) {
	t.Age = newAge
}

func Test_Defer(t *testing.T) {

	defer func() {
		fmt.Println("test1")
	}()

	fmt.Println("gap111")
	defer func() {
		fmt.Println("test2")
	}()
	fmt.Println("gap222")
	defer func() {
		fmt.Println("test3")
	}()

}

func Test_ReturnPointsOrAddr(t *testing.T) {

}

func getAddr(str string) *string {
	return &str
}

//func getPoints(str string) string {
//	return *str
//}
