package main

import "fmt"

func preOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.value)
	fmt.Print("->")
	preOrderTraverse(root.left)
	preOrderTraverse(root.right)
}
