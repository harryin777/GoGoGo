package main

import "fmt"

// https://dream.blog.csdn.net/article/details/128989981
// 字符串加密
func main22() {
	var count int
	fmt.Scan(&count)
	data := make([]string, 0, count)
	for c := 0; c < count; c++ {
		var tmp string
		fmt.Scan(&tmp)
		data = append(data, tmp)
	}
	cal22(data)
}

func cal22(str []string) {
	for i := 0; i < len(str); i++ {
		ans := ""
		for j := 0; j < len(str[i]); j++ {
			if str[i][j]+uint8(pos(j)) > 'z' {
				ans += string('a' + str[i][j] + uint8(pos(j)) - 'z' - 1)
			} else {
				ans += string(str[i][j] + uint8(pos(j)))
			}
		}
		fmt.Println(ans)
	}
}

func pos(p int) int {
	arr := []int{1, 2, 4}
	if p < 3 {
		return arr[p]
	}

	count := 3
	var ans int
	tmp1 := arr[0]
	tmp2 := arr[1]
	tmp3 := arr[2]
	for count <= p {
		ans = tmp1 + tmp2 + tmp3
		tmp1 = tmp2
		tmp2 = tmp3
		tmp3 = ans
		count++
	}

	return ans
}
