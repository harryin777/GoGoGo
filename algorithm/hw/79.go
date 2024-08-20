package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
*计算三叉搜索树的高度
 */

type tt struct {
	left  *tt
	mid   *tt
	right *tt
	val   int
}

func main79() {
	var count int
	fmt.Scanln(&count)
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	arr := strings.Split(buf.Text(), " ")
	nums := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		val, _ := strconv.Atoi(arr[i])
		nums = append(nums, val)
	}
	cal79(nums)
}

func cal79(nums []int) {
	root := &tt{
		val: nums[0],
	}
	var dfs func(*tt, int) *tt
	dfs = func(root *tt, num int) *tt {
		if root == nil {
			root = &tt{
				val: num,
			}
			return root
		}
		if root.val-500 > num {
			root.left = dfs(root.left, num)
		} else if root.val+500 < num {
			root.right = dfs(root.right, num)
		} else {
			root.mid = dfs(root.mid, num)
		}
		return root
	}
	for i := 1; i < len(nums); i++ {
		dfs(root, nums[i])
	}
	fmt.Println(calttHeight(root))
}

func calttHeight(root *tt) int {
	if root == nil {
		return 0
	}

	return max(max(calttHeight(root.left), calttHeight(root.mid)), calttHeight(root.right)) + 1
}
