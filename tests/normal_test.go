package tests

//标准输出包
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
	Male bool
}

/**
 * @Description 测试遍历字符串
 * @Param
 * @return
 **/
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

/**
 * @Description 测试数组
 * @Param
 * @return
 **/
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

/**
 * @Description 测试slice
 * @Param
 * @return
 **/
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
 * @Description make(map[KeyType]ValueType, [cap])
 * @Param
 * @return
 **/
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

	//结构体数组
	tmpArray := []tem{
		{
			Name: "1",
			Age:  1,
		},
		{
			Name: "2",
			Age:  2,
		},
	}
	_ = tmpArray
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

/**
 * @Description 测试defer
 * @Param
 * @return
 **/
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

/**
 * @Description 序列化测试，唯一的就是golang可以把字节数组转换成json格式的字符串，java并不行
				因为golang的字符串是有字节组成的，并不像传统的字符串是由字符组成的
 * @Param
 * @return
 **/
func Test_Str2JSON(t *testing.T) {
	str := "[\n\t2,\n\t3\n]"
	_ = str
	data := []int{2, 3}
	var result []int
	//strBytes 是一个字节数组
	strBytes, _ := json.Marshal(data)
	println(strBytes)
	//这里输出的把字节数组转换成了一个str
	println(string(strBytes))
	json.Unmarshal([]byte(strBytes), &result)
	for _, v := range result {
		println(v)
	}

	//strBytes2, _ := json.Marshal(str)
	//println(string(strBytes2))
}

/**
 * @Description  测试高效字符串拼接
 * @Param
 * @return
 **/
func Test_BiteBuffer(t *testing.T) {
	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString("123")
	}

	fmt.Println(b.String())
}

/**
* @Description: 测试int类型
* @Param:
* @return:
**/
func Test_Int(t *testing.T) {
	//字符是放在int里，当然也可以是byte范围0-255
	var a int = '世'
	fmt.Println(fmt.Sprintf("原字符：%c, 对应码值：%d", a, a))
	fmt.Println("a 占用的空间 = ", unsafe.Sizeof(a), "字节")

}

/**
 * @Description 判断结构体是否为空
 * @Param
 * @return
 **/
func Test_StructIsEmptyOrNot(t *testing.T) {
	var temp tem
	if (temp == tem{}) {
		fmt.Println("为空")
	}
}

/**
 * @Description 测试空slice可不可以分割
 * @Param
 * @return
 **/
func Test_EmptySlice(t *testing.T) {
	var slice []int = make([]int, 0)
	slice1 := slice[0:1]
	fmt.Println(slice1)
}

func Test_test(t *testing.T) {
	/**
	var n1 int32 = 12
	var n2 int64
	var n3 int8
	这里是编译不通过的，因为 n1 +32 还是一个int32 golang必须显示类型转换，不同于java会自动转换类型
	n2 = n1 + 32
	n3 = n1 +2
	*/
	var a int32 = 12
	str := fmt.Sprintf("%d", a)
	fmt.Println(str)

	//prec是保留小数位，fmt是一种格式
	str2 := strconv.FormatFloat(float64(a), 'f', 10, 64)
	fmt.Println(str2)
}
