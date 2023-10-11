package embedStruct

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func Test_1(t *testing.T) {
	b := C{
		B{
			"aaa",
		},
	}
	ans, err := jsoniter.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n ", ans)

	b2 := D{
		B: B{
			Str: "aaa",
		},
	}
	ans2, err := jsoniter.Marshal(b2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n ", ans2)
}
