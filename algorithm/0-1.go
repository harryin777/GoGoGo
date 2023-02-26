package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := strings.Split(sc.Text(), " ")
	money, _ := strconv.Atoi(str[0])
	n, _ := strconv.Atoi(str[1])
	items := make([]item, n)
	for i := 0; i < n; i++ { //构造物品
		sc.Scan()
		it := strings.Split(sc.Text(), " ")
		v, _ := strconv.Atoi(it[0])
		p, _ := strconv.Atoi(it[1])
		q, _ := strconv.Atoi(it[2])
		items[i] = item{
			v,
			p,
			q,
			nil,
			nil,
		}

	}

	for i := 0; i < len(items); i++ {
		var tmp = i
		if items[i].q == 0 {
			continue
		}
		if items[items[i].q-1].acc1 == nil {
			items[items[i].q-1].acc1 = &items[tmp]
		} else if items[items[i].q-1].acc2 == nil {
			items[items[i].q-1].acc2 = &items[tmp]
		}
	}

	dp := make([][]int, len(items)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, money+1)
	}

	cnt := 1
	for i := 0; i < len(items); i++ {
		if items[i].q != 0 {
			continue
		}

		var (
			mainP    = items[i].v
			mainMYD  = items[i].p * items[i].v
			main1P   = mainP
			main1MYD = mainMYD
			main2P   = mainP
			main2MYD = mainMYD
			main3P   = mainP
			main3MYD = mainMYD
		)
		if items[i].acc1 != nil {
			main1P += items[i].acc1.v
			main1MYD += items[i].acc1.p * items[i].acc1.v
		}
		if items[i].acc2 != nil {
			main2P += items[i].acc2.v
			main2MYD += items[i].acc2.p * items[i].acc2.v
		}
		if items[i].acc2 != nil && items[i].acc1 != nil {
			main3P += items[i].acc2.v + items[i].acc1.v
			main3MYD += items[i].acc1.p*items[i].acc1.v + items[i].acc2.p*items[i].acc2.v
		}

		for j := 1; j <= money; j++ {
			dp[cnt][j] = dp[cnt-1][j]
			if j >= mainP {
				dp[cnt][j] = max(dp[cnt][j], dp[cnt-1][j-mainP]+mainMYD)
			}
			if j >= main1P && main1P > mainP {
				dp[cnt][j] = max(dp[cnt][j], dp[cnt-1][j-main1P]+main1MYD)
			}
			if j >= main2P && main2P > mainP {
				dp[cnt][j] = max(dp[cnt][j], dp[cnt-1][j-main2P]+main2MYD)
			}
			if j >= main3P && main3P > mainP {
				dp[cnt][j] = max(dp[cnt][j], dp[cnt-1][j-main3P]+main3MYD)
			}

		}
		cnt++
	}
	fmt.Println(dp[cnt-1][money])
}

type item struct {
	v    int
	p    int
	q    int
	acc1 *item
	acc2 *item
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
