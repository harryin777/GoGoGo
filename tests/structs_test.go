/**
*   @Author: yky
*   @File: structs_test
*   @Version: 1.0
*   @Date: 2021-05-27 22:15
 */
package tests

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/goinggo/mapstructure"
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
	fmt.Printf("map2struct后得到的 struct 内容为:%v", t1)

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
