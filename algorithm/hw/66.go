package main

import (
	"fmt"
)

/*
*
橱窗里有一排宝石，不同的宝石对应不同的价格，宝石的价格标记为gems[i],
宝石可同时出售1个或多个，如果同时出售多个，则要求出售的宝石编号连续；
例如客户最大购买宝石个数为/购买的宝石编号必须为·
拥有总面值为的钱，请问最多能购买到多少个宝石。如无法购买宝石，则返回。
输入描述

第一行输入参数类型为int，取值范围：1驴6]/表示橱窗中宝石的总数量。
之后行分另刂表示从第个到第一1个宝石的价格，即到gems[n-l]的价格，类型为int/取值范围：
之后一行输入v/类型为int，取值范围：19^9]表示你拥有的钱。
输出描述
输出int类型的返回值，表示最大可购买的宝石数量。
7
8
4
6
3
1
6
7
10

3

7
8
4
6
3
1
1
7
12

4
*/
func main66() {
	var gemCount int
	fmt.Scanln(&gemCount)
	var gemVals []int
	for i := 0; i < gemCount; i++ {
		var val int
		fmt.Scan(&val)
		gemVals = append(gemVals, val)
	}
	var totalMoney int
	fmt.Scanln(&totalMoney)
	cal66(gemCount, totalMoney, gemVals)
}

func cal66(gemCount, totalMoney int, gemVals []int) {
	i, j, ans := 0, 0, 0
	curMoney := totalMoney
	for j < gemCount {
		if curMoney >= gemVals[j] {
			curMoney -= gemVals[j]
			j++
		} else {
			ans = max(ans, j-i)
			curMoney = totalMoney
			i++
			j = i
		}
	}

	fmt.Println(ans)
}
