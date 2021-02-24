package tree

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
* inorder: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/submissions/
* preorder: https://leetcode-cn.com/problems/binary-tree-preorder-traversal/submissions/
* postorder: https://leetcode-cn.com/problems/binary-tree-postorder-traversal/submissions/
**/

// 递归方式
func inorderD(root *TreeNode) []int {
	// 左 -> 根 -> 右
	res := []int{}

	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return res
}

func preorderD(root *TreeNode) []int {
	// 根 -> 左 -> 右
	res := []int{}

	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	preorder(root)

	return res
}

func postorderD(root *TreeNode) []int {
	// 左 -> 右 -> 根
	res := []int{}

	var postorder func(root *TreeNode)
	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)
		postorder(root.Right)
		res = append(res, root.Val)
	}

	postorder(root)
	return res
}

// 迭代方式

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	count := queue.Len()

	for count != 0 {
		levelRes := []int{}

		for i := 0; i < count; i++ {
			item := queue.Front()
			queue.Remove(item)

			node := item.Value.(*TreeNode)
			levelRes = append(levelRes, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		count = queue.Len()
		res = append(res, levelRes)
	}
	return res
}
