package tests

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	fmt.Println(math.Abs(float64(-1))) //取到绝对值
	fmt.Println(math.Ceil(3.8))        //向上取整
	fmt.Println(math.Floor(3.6))       //向下取整
	fmt.Println(math.Mod(11, 3))       //取余数 11%3 效果一样
	fmt.Println(math.Modf(3.22))       //取整数跟小数
	fmt.Println(math.Pow(3, 2))        //X 的 Y次方  乘方
	fmt.Println(math.Pow10(3))         //10的N次方 乘方
	fmt.Println(math.Sqrt(9))          //开平方  3
	fmt.Println(math.Cbrt(8))          //开立方  2
	fmt.Println(math.Pi)               //π
	fmt.Println(math.Round(4.2))       //四舍五入

	fmt.Println(math.IsNaN(3.4))      //false   报告f是否表示一个NaN（Not A Number）值。
	fmt.Println(math.Trunc(1.999999)) //1    返回整数部分（的浮点值）。
	fmt.Println(math.Max(-1.3, 0))    //0   返回x和y中最大值
	fmt.Println(math.Min(-1.3, 0))    //-1.3  返回x和y中最小值
	fmt.Println(math.Dim(-12, -19))   //7 函数返回x-y和0中的最大值
	fmt.Println(math.Dim(-12, 19))    //0 函数返回x-y和0中的最大值
	fmt.Println(math.Cbrt(8))         //2  返回x的三次方根
	fmt.Println(math.Hypot(3, 4))     //5  返回Sqrt(p*p + q*q)
	fmt.Println(math.Pow(2, 8))       //256  返回x^y
}
