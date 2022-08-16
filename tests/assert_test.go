package tests

import (
	"fmt"
	"os"
	"testing"
)

func VarType(s interface{}) (err error) {

	switch s.(type) {
	case int:
		fmt.Println("这是个 int")
		fmt.Println(s)
	case string:
		fmt.Println("这是个 string")
		fmt.Println(s)
	case int8:
		fmt.Println("这是个 int8")
		fmt.Println(s)
	default:
		fmt.Println("这是个啥")
		fmt.Printf("%f : %v", s, s)
	}

	if sp, ok := s.(string); ok {
		fmt.Println("果然是个 string")
		fmt.Println(sp)
	}
	return nil
}

/**
 * @Description 断言的用法
 * @Param
 * @return
 **/
func Test_Assert(t *testing.T) {
	_ = VarType("oiu")
}

func Test_exist(t *testing.T) {
	defer fmt.Println(12)
	os.Exit(3)
}
