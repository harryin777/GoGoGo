package binaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// pre 前序
func pre(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	order := []int{root.Val}
	//if root.Left == nil && root.Right == nil {
	//
	//}
	if root.Left != nil {
		order = append(order, pre(root.Left)...)
	}

	if root.Right != nil {
		order = append(order, pre(root.Right)...)
	}

	return order
}

// mid 中序
func mid(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	order := []int{}

	if root.Left != nil {
		order = append(order, pre(root.Left)...)
	}

	order = append(order, root.Val)

	if root.Right != nil {
		order = append(order, pre(root.Right)...)
	}

	return order
}

// after 后序
func after(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	order := []int{}

	if root.Left != nil {
		order = append(order, pre(root.Left)...)
	}

	if root.Right != nil {
		order = append(order, pre(root.Right)...)
	}

	order = append(order, root.Val)

	return order
}
