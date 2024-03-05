package tests

//标准输出包
import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"math"
	"math/rand"
	url2 "net/url"
	"regexp"
	"strconv"
	"testing"
	"time"
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
 * @Description 测试空slice可不可以分割
 * @Param
 * @return
 **/
func Test_EmptySlice(t *testing.T) {
	//var slice []int = make([]int, 0)
	//slice1 := slice[0:1]
	//fmt.Println(slice1)
	slice2 := []string{
		"1", "2", "3",
	}
	fmt.Println(slice2[0:4])
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
	fmt.Printf("type is %T \n", str)

	//prec是保留小数位，fmt是一种格式
	str2 := strconv.FormatFloat(float64(a), 'f', 10, 64)
	fmt.Println(str2)
}

/**
* @Description: 匿名函数，可以用在函数体内部，函数体内部不能声明带有名字的函数，匿名函数是一种解决方案
* @Param:
* @return:
**/
var f1 = func(x, y int) {
	fmt.Printf("anonymous func p1:%d, p2:%d \n", x, y)
}

func Test_AnonymousFunc(t *testing.T) {
	f1(1, 2)
	//只执行一次的函数
	func() {
		fmt.Println("立即执行的函数！")
	}()
}

/**
 * @Description 正则的使用
 * @Param
 * @return
 **/
func Test_Regex(t *testing.T) {
	//判断一个字符串是否为纯数字，小数不行
	//[0-9]* 这个也行
	matched, err := regexp.Match("^[-+]?[\\d]*$", []byte("-10.00"))
	fmt.Println(matched, err) //true nil
}

/**
 * @Description 转义函数
 * @Param
 * @return
 **/
func Test_urlPathUnescape(t *testing.T) {
	url := "http://sdf\nsdf\nqweqwe"
	//url被转义了
	fmt.Println(url2.PathUnescape(url))
	//url没有被转义
	fmt.Println(url2.PathEscape(url))
}

/**
 * @Description 初始化结构体数组
 * @Param
 * @return
 **/
func Test_initStructArray(t *testing.T) {
	type a1 struct {
		Name string `json:"name"`
	}

	a1s := []a1{
		a1{
			"a1",
		},
		a1{
			"a2",
		},
	}

	_ = a1s
}

func Test_slice2(t *testing.T) {
	a := []int{1, 2, 3}
	//a = append(a[:0], a[1:]...) // 删除开头1个元素
	//这是清整个 slice 的操作
	a = a[:0]
	fmt.Println(len(a))
}

/*
*

	 @Description 合并两个 slice
	 @Param
	 @return
	*
*/
func Test_Merge2Slices(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 6, 7}
	a = append(a, b...) // 删除开头1个元素
	fmt.Println(len(a))
}

/*
*

	 @Description rand 随机数
	 @Param
	 @return
	*
*/
func Test_Rand(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		fmt.Println(r.Intn(200))
	}

}

/**
 * @Description json 多传一个参数在序列化的时候,如果不用到,不会有什么影响
 * @Param
 * @return
 **/
func Test_Json(t *testing.T) {
	type A struct {
		AA string `json:"aa"`
		AB string `json:"ab"`
	}

	type a struct {
		AA string `json:"aa"`
	}

	A1 := A{
		"qq",
		"ww",
	}
	b, _ := jsoniter.Marshal(A1)
	var a1 a
	jsoniter.Unmarshal(b, &a1)
	fmt.Println(a1)
}

/**
 * @Description json 序列化之后,一个字符串的字节长度是多少,比原来多了还是少了
 * @Param
 * @return
 **/
func Test_Json2(t *testing.T) {
	str := "这是一个字符串"
	b, _ := jsoniter.Marshal(str)
	fmt.Printf("before json : %v \n", len(str))
	fmt.Printf("after json : %v \n", len(b))

	num := "657456"
	b2, _ := jsoniter.Marshal(num)
	fmt.Println(num)
	fmt.Printf("before json : %v \n", len(num))
	fmt.Println(b2)
	fmt.Printf("after json : %v \n", len(b2))

}

// 比较两个简单结构体，验证是不是值传递
func Test_CompareSimpleStruct(t *testing.T) {
	//type a struct {
	//	Name string `json:"name"`
	//}
	//a1 := a{
	//	"lalala",
	//}
	//b1 := a1
	//fmt.Println(a1)
	//fmt.Println(b1)
	//a1.Name = "qqqq"
	//fmt.Println(a1)
	//fmt.Println(b1)
	////普通的类型时值传递，负责类型，比如map和slice是指针传递
	//slice1 := []int{1, 2, 3}
	//slice2 := slice1
	//fmt.Println(slice2)
	//slice1[1] = 1
	//fmt.Println(slice2)
	var b float32
	b = 1.1
	fmt.Println(int(math.Ceil(float64(b))))

}

type Person2 struct {
	name  string
	hobby string
}

func (p Person2) MarshalJSON() ([]byte, error) {
	return []byte(`{"name":"` + p.name + `","hobby":"` + p.hobby + `"}`), nil
}

func Test_Marshal(t *testing.T) {
	var i interface{}
	fmt.Println(i.(string))
}

func getPersonName(level int, p []Person) (res []string) {
	res = make([]string, 0, 10)
	name := p[level].Name
	res = append(res, name)
	res2 := getPersonName(level+1, p)
	res = append(res, res2...)
	return
}

func Test_hello2(t *testing.T) {
	//T2()
	fmt.Println(strconv.ParseFloat(fmt.Sprintf("%.4f", float64(2.55)/6.88), 64))
}

func T2() int {
	defer TTTT("defer")
	return TTTT("return")
}

func TTTT(str string) int {
	fmt.Printf("hello: %v \n", str)
	return 1
}
