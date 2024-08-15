package orderedMap

import (
	"fmt"
	"testing"
)

func Test_t1(t *testing.T) {
	om := NewOrderedMap[string, int]()
	om.Set("first", 1)
	om.Set("second", 2)
	om.Set("third", 3)

	fmt.Println("Keys:", om.Keys()) // [first second third]
	fmt.Println("Values:", om.Values())
	val, _ := om.Get("second")      // [1 2 3]
	fmt.Println("Get second:", val) // 2, true

	key, value, found := om.At(1)
	fmt.Printf("At index 1: %v = %v, found: %v\n", key, value, found) // second = 2, found: true

	om.Delete("second")
	fmt.Println("Keys after delete:", om.Keys()) // [first third]
}
