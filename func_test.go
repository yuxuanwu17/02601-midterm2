package main

import (
	"fmt"
	"testing"
)

func TestIsBSTOrdered(t *testing.T) {
	// [2,1,3]
	root := TreeNode{
		left: &TreeNode{
			left:  nil,
			right: nil,
			value: 1,
		},
		right: &TreeNode{
			left:  nil,
			right: nil,
			value: 3,
		},
		value: 2,
	}

	fmt.Println(IsBSTOrdered(&root))

	//root = [5,1,4,null,null,3,6]
	node3 := TreeNode{
		left:  nil,
		right: nil,
		value: 3,
	}

	node6 := TreeNode{
		left:  nil,
		right: nil,
		value: 6,
	}
	node4 := TreeNode{
		left:  &node3,
		right: &node6,
		value: 4,
	}
	node1 := TreeNode{
		left:  nil,
		right: nil,
		value: 1,
	}
	node5 := TreeNode{
		left:  &node1,
		right: &node4,
		value: 5,
	}

	fmt.Println(IsBSTOrdered(&node5))
}
