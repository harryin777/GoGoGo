package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/**
题目描述
有一个考古学家发现一个石碑，但是很可惜发现时其已经断成多段，原地发现N个断口整齐的石碑碎片，为了破解石碑
内容
考古学家希望有程序能帮忙计算复原后的石碑文字，你能帮忙吗
备注
如果存在石碑碎片内容完全相同，则由于碎片间的顺序不影响复原后的碑文内容，
输入描述
输出描述
仅相同碎片间的亻立置变化不影响组合
第一行输入，N表示石碑碎片的个数第二行依次输入石碑碎片上的文字内容s共有组
输出石碑文字的组合（按照升序排列），行尾无多余空格
*/

func main70() {
	buf := bufio.NewScanner(os.Stdin)
	buf.Scan()
	n := buf.Text()
	nInt, _ := strconv.Atoi(n)
	buf.Scan()
	str := buf.Text()
	arr := strings.Split(str, " ")
	cal70(nInt, arr)
}

func cal70(n int, arr []string) {
	removeDup := make(map[string]struct{})
	total := make([]string, 0, 10)
	var dfs func([]string, map[string]struct{})
	dfs = func(cur []string, exi map[string]struct{}) {
		if len(cur) == len(arr) {
			total = append(total, strings.Join(cur, ""))
			return
		}
		for i := 0; i < len(arr); i++ {
			if _, ok := exi[arr[i]]; ok {
				continue
			}
			exi[arr[i]] = struct{}{}
			cur = append(cur, arr[i])
			dfs(cur, exi)
			cur = cur[:len(cur)-1]
			delete(exi, arr[i])
		}
	}
	exi := make(map[string]struct{})
	dfs([]string{}, exi)
	for _, val := range total {
		removeDup[val] = struct{}{}
	}
	ans := make([]string, 0, len(total))
	for key, _ := range removeDup {
		ans = append(ans, key)
	}

	sort.Strings(ans)

	fmt.Println(ans)
}
