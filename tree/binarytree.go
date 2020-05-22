package tree

//链表存储的二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var stack []*TreeNode
	var res []int
	stack = append(stack, root)
	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Val) //打印根
		if e.Right != nil {      //将右边压入栈
			stack = append(stack, e.Right)
		}
		if e.Left != nil { //将左边压入栈
			stack = append(stack, e.Left)
		}
	}
	return res
}

//中序遍历
func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := inOrderTraversal(root.Left) //先打印左半边
	res = append(res, root.Val)
	res = append(res, inOrderTraversal(root.Right)...) //再打印右半边
	return res
}

//后序排序
func postOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		lref := postOrderTraversal(root.Left)
		if len(lref) > 0 {
			res = append(res, lref...)
		}
	}
	if root.Right != nil {
		rref := postOrderTraversal(root.Right)
		if len(rref) > 0 {
			res = append(res, rref...)
		}
	}
	res = append(res, root.Val)
	return res
}
