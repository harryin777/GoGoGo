package binaryTree

import (
	"fmt"
	"testing"
)

//
//
var r TreeNode = TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: nil,
	},
}

func TestBinaryTreePre(t *testing.T) {
	fmt.Println(pre(&r))
	fmt.Println(mid(&r))
	fmt.Println(after(&r))
}
