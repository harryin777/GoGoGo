/**
*   @Author: yky
*   @File: reflect_test
*   @Version: 1.0
*   @Date: 2021-07-05 21:39
 */
package tests

import (
	"fmt"
	"reflect"
	"testing"
)

type Drawer struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/**
* @Description: 简单获取普通结构体的属性和属性值
* @Param:
* @return:
**/
func Test_GetAttri(t *testing.T) {
	testMap := make(map[string]Drawer)
	testDrawer := Drawer{
		Name: "test",
		Age:  12,
	}
	testMap["persion1"] = testDrawer

	testStr := "strstrstr"

	for _, val := range testMap {
		dataType := reflect.TypeOf(val)
		fmt.Printf("%v \n", dataType.Field(0))
		fmt.Printf("%v \n", dataType.Field(1))
		dataVal := reflect.ValueOf(val)
		fmt.Printf("%v \n", dataVal.Field(0))
		fmt.Printf("%v \n", dataVal.Field(1))
		val.Age = 15
	}
	testDrawer.Age = 18
	fmt.Println(reflect.ValueOf(testDrawer))

	fmt.Println(testMap)

	fmt.Println("------------------------")
	fmt.Println(reflect.ValueOf(&testMap).Elem())
	fmt.Printf("type : %T \n", reflect.ValueOf(&testMap).Elem())
	fmt.Println(reflect.ValueOf(testMap))
	fmt.Printf("type : %T \n", reflect.ValueOf(testMap))
	fmt.Println(reflect.Indirect(reflect.ValueOf(testMap)))
	fmt.Printf("type : %T \n", reflect.Indirect(reflect.ValueOf(testMap)))
	fmt.Println("------------------------")
	fmt.Println(reflect.ValueOf(&testStr).Elem())
	fmt.Printf("type : %T \n", reflect.ValueOf(&testStr).Elem())
	fmt.Println(reflect.ValueOf(&testStr))
	fmt.Printf("type : %T \n", reflect.ValueOf(&testStr))
	fmt.Println(reflect.Indirect(reflect.ValueOf(&testStr)))
	fmt.Printf("type : %T \n", reflect.Indirect(reflect.ValueOf(&testStr)))
	fmt.Println("------------------------")
	fmt.Println(reflect.ValueOf(&testDrawer).Elem())
	fmt.Printf("type : %T \n", reflect.ValueOf(&testDrawer).Elem())
	fmt.Println(reflect.ValueOf(&testDrawer))
	fmt.Printf("type : %T \n", reflect.ValueOf(&testDrawer))
	fmt.Println(reflect.Indirect(reflect.ValueOf(&testDrawer)))
	fmt.Printf("type : %T \n", reflect.Indirect(reflect.ValueOf(&testDrawer)))
}

type CrazyDrawer struct {
	Name        string        `json:"name"`
	Age         int           `json:"age"`
	MasterPiece []MasterPiece `json:"masterPiece"`
}

type MasterPiece struct {
	Id    int    `json:"id"`
	MName string `json:"mName"`
}

// 在 map 中修改结构体的值然后给 map 重新赋值
func Test_InnerSliceOfStruct(t *testing.T) {
	masterPieces := []MasterPiece{
		{
			Id:    1,
			MName: "m1",
		},
		{
			Id:    2,
			MName: "m2",
		},
	}
	crazyOne := CrazyDrawer{
		Name:        "crazy",
		Age:         22,
		MasterPiece: masterPieces,
	}
	cmap := make(map[string]CrazyDrawer)
	cmap["c1"] = crazyOne

	for _, val := range cmap {
		dataVal := reflect.ValueOf(val)
		//d2 := reflect.ValueOf(val).Elem()
		//_ = d2
		fmt.Printf("%v \n", dataVal.Field(0))
		fmt.Printf("%v \n", dataVal.Field(1))
		fmt.Printf("%v \n", dataVal.Field(2))
		//如何获取结构体中结构体数组的属性值
		//dataVal2 := dataVal.Field(2)
		//fmt.Printf("%v \n", dataVal2)

		dataVal1 := reflect.ValueOf(&val).Elem()
		////修改属性值
		dataVal1.FieldByName("Name").SetString("newName")
		cmap["c1"] = dataVal1.Interface().(CrazyDrawer)
		fmt.Printf("%v \n", dataVal1.Field(0))
	}

	fmt.Println(cmap)
}

// 通过反射修改变量的值
func Test_ReflectNormal(t *testing.T) {
	// 声明整型变量a并赋初值
	var a int = 1024
	// 获取变量a的反射值对象(a的地址)
	valueOfA := reflect.ValueOf(&a)
	// 取出a地址的元素(a的值)
	valueOfA = valueOfA.Elem()
	// 修改a的值为1
	valueOfA.SetInt(1)
	// 打印a的值
	fmt.Println(valueOfA.Int())
	fmt.Println(a)
}

// 反射遍历所有属性
func Test_Traverse(t *testing.T) {
	// 定义一个结构体
	p := Person{
		Name: "Alice",
		Age:  20,
	}

	// 获取结构体类型
	t1 := reflect.TypeOf(p)

	// 遍历结构体的所有属性
	values := make(map[string]interface{})
	for i := 0; i < t1.NumField(); i++ {
		// 获取属性名
		name := t1.Field(i).Name

		// 如果属性名与给定字符串相等
		if name == "Age" {
			// 获取属性值
			value := reflect.ValueOf(p).Field(i).Interface()

			// 将属性值放入 map 中
			values[name] = value
		}
	}

	// 输出结果
	fmt.Println(values)
}
