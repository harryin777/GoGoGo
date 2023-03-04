package main

import "fmt"

// https://dream.blog.csdn.net/article/details/129019183
// 二叉树的层次遍历
func main38() {
	var last, mid string
	fmt.Scan(&last, &mid)
	cal38(last, mid)

}

func cal38(last string, mid string) {
	node := rebuild(mid, last)
	queue := make([]*Node38, 0, 10)
	queue = append(queue, node)
	for len(queue) != 0 {
		lll := len(queue)
		tmp := make([]*Node38, 0, 10)

		for i := 0; i < lll; i++ {
			one := queue[0]
			queue = queue[1:]
			tmp = append(tmp, one)
			if one.Left != nil {
				queue = append(queue, one.Left)
			}
			if one.Right != nil {
				queue = append(queue, one.Right)
			}
		}
		for i := 0; i < len(tmp); i++ {
			fmt.Printf("%v", tmp[i].Val)
		}
	}

}

func rebuild(inorder string, postorder string) *Node38 {
	num := len(postorder)
	if num == 0 {
		return nil
	}
	val := postorder[num-1]
	var k int
	var v rune
	for k, v = range inorder {
		if rune(val) == v {
			break
		}
	}
	node := new(Node38)
	node.Val = string(val)
	inorder1 := inorder[:k]
	inorder2 := inorder[k+1:]
	postorder1 := postorder[:k]
	postorder2 := postorder[k : len(postorder)-1]
	node.Left = rebuild(inorder1, postorder1)
	node.Right = rebuild(inorder2, postorder2)
	return node
}

type Node38 struct {
	Val   string
	Left  *Node38
	Right *Node38
}
