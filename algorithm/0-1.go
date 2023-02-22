package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	c := 0
	count := 0
	sum, total := 0, 0
	GoodsList := make([]Goods, 0, 100)
	for {

		line, _ := fmt.Scanln(&a, &b, &c)
		if line == 0 {
			break
		}

		if count == 0 {
			sum = a
			total = b
		} else {
			GoodsList = append(GoodsList, Goods{
				Price: a,
				Z:     b,
				P:     c,
			})
		}
		count++
	}

	mainPro := 0
	for index, goods := range GoodsList {
		if goods.P == 0 {
			mainPro++
			continue
		}

		if GoodsList[goods.P-1].Acc1 == nil {
			GoodsList[goods.P-1].Acc1 = &GoodsList[index]
		} else if GoodsList[goods.P-1].Acc2 == nil {
			GoodsList[goods.P-1].Acc2 = &GoodsList[index]
		}
	}

	dp := make([][]int, total+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum+1)
	}

	for i := 1; i <= total; i++ {
		if GoodsList[i-1].P != 0 {
			continue
		}
		var (
			mainMYD    = GoodsList[i-1].Price * GoodsList[i-1].Z
			mainPrice  = GoodsList[i-1].Price
			mainMYD1   = mainMYD
			mainPrice1 = mainPrice
			mainMYD2   = mainMYD
			mainPrice2 = mainPrice
			mainMYD3   = mainMYD
			mainPrice3 = mainPrice
		)
		if GoodsList[i-1].Acc1 != nil {
			mainMYD1 += GoodsList[i-1].Acc1.Price * GoodsList[i-1].Acc1.Z
			mainPrice1 += GoodsList[i-1].Acc1.Price
		}
		if GoodsList[i-1].Acc2 != nil {
			mainMYD2 += GoodsList[i-1].Acc2.Price * GoodsList[i-1].Acc2.Z
			mainPrice2 += GoodsList[i-1].Acc2.Price
		}
		if GoodsList[i-1].Acc2 != nil && GoodsList[i-1].Acc1 != nil {
			mainMYD3 += GoodsList[i-1].Acc2.Price*GoodsList[i-1].Acc2.Z + GoodsList[i-1].Acc1.Price*GoodsList[i-1].Acc1.Z
			mainPrice3 += GoodsList[i-1].Acc2.Price + GoodsList[i-1].Acc1.Price
		}

		fmt.Println("gz", mainPrice, mainMYD, mainPrice1, mainMYD1, mainPrice2, mainMYD2, mainPrice3, mainMYD3)

		for j := 1; j <= sum; j++ {
			if j >= mainPrice {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-mainPrice]+mainMYD)
			}
			if j >= mainPrice1 && mainPrice1 > mainPrice {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-mainPrice1]+mainMYD1)
			}
			if j >= mainPrice2 && mainPrice2 > mainPrice {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-mainPrice2]+mainMYD2)
			}
			if j >= mainPrice3 && mainPrice3 > mainPrice {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-mainPrice3]+mainMYD3)
			}
		}
	}

	fmt.Println(dp[total][sum])
}

type Goods struct {
	Price int
	Z     int
	P     int
	Acc1  *Goods
	Acc2  *Goods
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
