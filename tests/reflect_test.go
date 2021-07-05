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
	testMap := make(map[string]interface{})
	testMap["persion1"] = Drawer{
		Name: "abc",
		Age:  12,
	}

	for _, val := range testMap {
		dataType := reflect.TypeOf(val)
		fmt.Printf("%v \n", dataType.Field(0))
		fmt.Printf("%v \n", dataType.Field(1))
		dataVal := reflect.ValueOf(val)
		fmt.Printf("%v \n", dataVal.Field(0))
		fmt.Printf("%v \n", dataVal.Field(1))
	}
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
	cmap := make(map[string]interface{})
	cmap["c1"] = crazyOne

	for _, val := range cmap {
		dataVal := reflect.ValueOf(val)
		fmt.Printf("%v \n", dataVal.Field(0))
		fmt.Printf("%v \n", dataVal.Field(1))
		fmt.Printf("%v \n", dataVal.Field(2))
		//如何获取结构体中结构体数组的属性值
		dataVal2 := dataVal.Field(2)

		fmt.Printf("%v \n", dataVal2)

	}
}
