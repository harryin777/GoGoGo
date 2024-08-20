package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
黑白图像常采用灰度图的方式存储，即图像的每个像素填充一个灰阶值，
256阶灰度图是一个灰阶值取值范围为0-255的灰阶矩阵，0表示全黑，255表示全白，
范围内的其他值表示不同的灰度，比如下面的图像及其对应的灰阶矩阵:

 但在计算机中实际存储时，会使用压缩算法，其中一种压缩格式和描述如下：

10 10 255 34 0 1 255 8 0 3 255 6 0 5 255 4 0 7 255 2 0 9 255 2 1

1、所有数值以空格分隔
2、前两个数分别表示矩阵的行数和列数
3、从第三个数开始，每两个数一组，每组第一个数是灰阶值，第二个数表示该灰阶值从左到右，
从上到下(可理解为将二维数组按行存储在一维矩阵中)的连续像素个数。比如题目所述例子，“255 34"表示有连续34个像素的灰阶值是255。
如此，图像软件在打开此格式灰度图的时候，就可以根据此算法从压缩数据恢复出原始灰度图矩阵。请从输入的压缩数恢复灰度图原始矩阵，并返回指定像素的灰阶值。

输入描述:
10 10 255 34 0 1 255 8 0 3 255 6 0 5 255 4 0 7 255 2 0 9 255 2 1
3 4

10 10 56 34 99 1 87 8 99 3 255 6 99 5 255 4 99 7 255 2 99 9 255 21
3 4

*/

func main63() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputLine := scanner.Text()

	scanner.Scan()
	var x, y int
	// Sscanf 是从标准输入中解析
	fmt.Sscanf(scanner.Text(), "%d %d", &x, &y)
	// Scanf 是从标准输入中读取
	//fmt.Scanf("%v %v", &x, &y)
	cal63(inputLine, x, y)
}

func cal63(input string, x, y int) {
	strArr := strings.Split(input, " ")
	arr := make([]int, 0, len(strArr))
	for _, val := range strArr {
		valInt, _ := strconv.Atoi(val)
		arr = append(arr, valInt)
	}
	matrix := make([][]int, arr[0])
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, arr[1])
	}
	target := x*arr[1] + y + 1
	for i := 2; i < len(arr)-1; i = i + 2 {
		target -= arr[i+1]
		if target <= 0 {
			fmt.Println(arr[i])
			break
		}
	}
}
