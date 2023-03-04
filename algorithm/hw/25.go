package main

import "fmt"

// https://dream.blog.csdn.net/article/details/128980730
// 人数最多的站点
func main25() {
	var people int
	fmt.Scan(&people)
	list := make(map[station]struct{})
	var first, last int
	for i := 0; i < people; i++ {
		var begin, end int
		fmt.Scan(&begin, &end)
		first = min(first, begin)
		last = max(last, end)
		list[station{
			Begin: begin,
			End:   end,
		}] = struct{}{}
	}

	cal25(list, people, first, last)
}

func cal25(list map[station]struct{}, people, fir, las int) {
	currPeo := 0
	maxPeo := 0
	ans := 0
	queue := make(map[station]struct{})
	for i := fir; i <= las; i++ {
		for sta, _ := range queue {
			if sta.End < i {
				delete(queue, sta)
				currPeo--
			}
		}
		for sta, _ := range list {
			if sta.Begin == i {
				delete(list, sta)
				queue[sta] = struct{}{}
				currPeo++
				if currPeo > maxPeo {
					maxPeo = currPeo
					ans = i
				}
			}
		}
	}
	fmt.Println(ans)
}

type station struct {
	Begin int
	End   int
}
