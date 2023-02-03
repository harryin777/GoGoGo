/**
*   @Author: yky
*   @File: structs_test
*   @Version: 1.0
*   @Date: 2021-05-27 22:15
 */
package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/goinggo/mapstructure"
	"log"
	"os"
	"testing"
)

type temp struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/**
* @Description: 结构体转成map map转回结构体，需要两个包 "github.com/fatih/structs" "github.com/goinggo/mapstructure"
* @Param:
* @return:
**/
func Test_Struct2map(t *testing.T) {
	var t1 temp
	//把struct转换成了map，可以实现动态赋值，再来一个参数map，两个map双循环

	map1 := structs.Map(t1)
	for k, v := range map1 {
		fmt.Printf("before key:%v,val:%v \n", k, v)
	}

	map2 := map[string]interface{}{
		"Name": "qqq",
	}
	//或者循环参数map
	for k, v := range map2 {
		map1[k] = v
	}

	for k, v := range map1 {
		fmt.Printf(" after key:%v,val:%v \n", k, v)
	}

	var t2 temp
	_ = t2
	if err := mapstructure.Decode(map1, &t1); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("map2struct后得到的 construct 内容为:%v", t1)

}

/**
* @Description: json 2 map
* @Param:
* @return:
**/
func Test_Json2map(t *testing.T) {
	jsonStr := `
    {
        "name":"liangyongxing",
        "age":12
    }
    `
	var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println(err)
	}
	fmt.Println(mapResult)
}

/**
 * @Description 关于缩进的问题
 * @Param
 * @return
 **/
func Test_Indent(t *testing.T) {
	type Road struct {
		Name   string
		Number int
	}
	roads := []Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "=", "\t")
	out.WriteTo(os.Stdout)
}

var (
	mapGlobal = make([]int, 0, 2)
)

func getMapGlobal() {
	mapGlobal = append(mapGlobal, 1, 2)
}

func TestMap1(t *testing.T) {
	//map1 := make([]int, 0, 2)
	//getMapGlobal()
	//fmt.Printf("%p, %p \n", map1, mapGlobal)
	//map1 = mapGlobal

	//for key, val := range mapGlobal {
	//	map1[key] = val
	//}
	//fmt.Printf("%p, %p", map1, mapGlobal)
}
