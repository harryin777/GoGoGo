package tests

import (
	"testing"
	"unicode"
)

const (
	// 注意这个方向, 左右相反的方向,上下相反的方向中间都是隔了 1,后面代码是做了减法
	Left = iota
	Top
	Right
	Bottom
)

func TestRobotCoordinate(t *testing.T) {
	println(move("LFF", 0, 0, Top))
}

// x,y 是坐标 z 是朝向.
func move(cmd string, x0 int, y0 int, z0 int) (x, y, z int) {
	x, y, z = x0, y0, z0
	repeat := 0
	repeatCmd := ""
	for _, s := range cmd {
		switch {
		case unicode.IsNumber(s):
			// 注意这里解析字符串到对应的数字
			repeat = repeat*10 + (int(s) - '0')
		case s == ')':
			for i := 0; i < repeat; i++ {
				x, y, z = move(repeatCmd, x, y, z)
			}
			repeat = 0
			repeatCmd = ""
		case repeat > 0 && s != '(' && s != ')':
			repeatCmd = repeatCmd + string(s)
		case s == 'L':
			z = (z + 1) % 4
		case s == 'R':
			z = (z - 1 + 4) % 4
		case s == 'F':
			switch {
			case z == Left || z == Right:
				x = x - z + 1
			case z == Top || z == Bottom:
				y = y - z + 2
			}
		case s == 'B':
			switch {
			case z == Left || z == Right:
				x = x + z - 1
			case z == Top || z == Bottom:
				y = y + z - 2
			}
		}
	}
	return
}
