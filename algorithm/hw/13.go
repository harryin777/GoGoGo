package main

import (
	"fmt"
	"sort"
	"strings"
)

func main13() {
	var beginIndex, Count int
	fmt.Scan(&beginIndex, &Count)
	cList := make([]string, 0, Count)
	for i := 0; i < Count; i++ {
		var tmp string
		fmt.Scan(&tmp)
		cList = append(cList, tmp)
	}

	cal13(beginIndex, cList)
}

func cal13(b int, cList []string) {
	ans := ""
	qualifyList := make([]string, 0, len(cList))
	qualifyList = append(qualifyList, cList[b])
	for len(qualifyList) > 0 {
		one := qualifyList[0]
		ans = ans + one
		qualifyList = qualifyList[1:]
		var tmp []string
		for i := 0; i < len(cList); i++ {
			if i == b {
				continue
			}
			if strings.HasPrefix(cList[i], string(one[len(one)-1])) {
				qualifyList = append(qualifyList, cList[i])
			} else {
				tmp = append(tmp, cList[i])
			}
		}
		sort.Slice(qualifyList, func(i, j int) bool {
			if len(qualifyList[i]) > len(qualifyList[j]) {
				return true
			} else if len(qualifyList[i]) == len(qualifyList[j]) {
				for m := 0; m < len(qualifyList[i]); m++ {
					if qualifyList[i][m] < qualifyList[j][m] {
						return true
					} else if qualifyList[i][m] == qualifyList[j][m] {
						continue
					} else {
						return false
					}
				}
			} else {
				return false
			}

			return false
		})
		if len(qualifyList) > 0 {
			tmp = append(tmp, qualifyList[1:]...)
			cList = tmp
			qualifyList = qualifyList[0:1]
		}

	}

	fmt.Println(ans)
}
