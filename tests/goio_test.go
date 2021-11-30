/**
*   @Author: yky
*   @File: goio_test
*   @Version: 1.0
*   @Date: 2021-06-03 23:12
 */
package tests

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_GoIO(t *testing.T) {
	counts := make(map[string]int) //创建一个空map
	input_data := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		input_data.Scan()
		counts[input_data.Text()]++
	}
	for index, value := range counts { //通过range遍历map
		fmt.Printf("%s\t%d\n", index, value)
	}
}
