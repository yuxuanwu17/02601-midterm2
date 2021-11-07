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

func TestLowestCommonAncestor(t *testing.T) {
	root := TreeNode{
		value: 3,
		left: &TreeNode{
			value: 5,
			left:  &TreeNode{nil, nil, 6},
			right: &TreeNode{&TreeNode{nil, nil, 7}, &TreeNode{nil, nil, 4}, 2},
		},
		right: &TreeNode{
			value: 1,
			left:  &TreeNode{nil, nil, 0},
			right: &TreeNode{nil, nil, 8},
		},
	}

	// preorder traverse
	fmt.Println("The preorder traversal of the given graph")
	preOrderTraverse(&root)
	fmt.Println()
	fmt.Println("5,1 => 3")
	fmt.Println(LowestCommonAncestor(&root, 5, 1))
	fmt.Println("5,4 => 5")
	fmt.Println(LowestCommonAncestor(&root, 5, 4))
	fmt.Println("2,0=>3")
	fmt.Println(LowestCommonAncestor(&root, 2, 0))

}

func TestMergeSortedLists(t *testing.T) {
	l1 := ListNode{
		value: 1,
		next: &ListNode{
			value: 3,
			next: &ListNode{
				value: 6,
				next:  nil,
			},
		},
	}

	l2 := ListNode{
		value: 2,
		next: &ListNode{
			value: 2,
			next: &ListNode{
				value: 2,
				next: &ListNode{
					value: 2,
					next: &ListNode{
						value: 6,
						next:  nil,
					},
				},
			},
		},
	}

	mergedList := MergeSortedLists(&l1, &l2)
	for {
		if mergedList.next == nil {
			fmt.Println(mergedList.value)
			break
		}
		fmt.Println(mergedList.value)
		mergedList = mergedList.next
	}

	fmt.Println("==================")

	lnode6 := ListNode{
		next:  nil,
		value: 6,
	}
	lnode3 := ListNode{
		next:  &lnode6,
		value: 3,
	}
	lnode1 := ListNode{
		next:  &lnode3,
		value: 1,
	}
	lnode2 := ListNode{
		next:  &lnode6,
		value: 2,
	}
	lnode2_2 := ListNode{
		next:  &lnode2,
		value: 2,
	}
	lnode2_2_2 := ListNode{
		next:  &lnode2_2,
		value: 2,
	}
	lnode2_2_2_2 := ListNode{
		next:  &lnode2_2_2,
		value: 2,
	}
	//for {
	//	if lnode1.next == nil {
	//		fmt.Println(lnode1.value)
	//		break
	//	}
	//	fmt.Println(lnode1.value)
	//	lnode1 = *lnode1.next
	//}
	//
	//fmt.Println("==================")
	//for {
	//	if lnode2_2_2_2.next == nil {
	//		fmt.Println(lnode2_2_2_2.value)
	//		break
	//	}
	//	fmt.Println(lnode2_2_2_2.value)
	//	lnode2_2_2_2 = *lnode2_2_2_2.next
	//}

	mergedList2 := MergeSortedLists(&lnode1, &lnode2_2_2_2)
	for {
		if mergedList2.next == nil {
			fmt.Println(mergedList2.value)
			break
		}
		fmt.Println(mergedList2.value)
		mergedList2 = mergedList2.next
	}

}

func TestMthFromEnd(t *testing.T) {
	// [1,2,3,4,5]
	ll := ListNode{
		value: 1,
		next: &ListNode{
			value: 2,
			next: &ListNode{
				value: 3,
				next: &ListNode{
					value: 4,
					next: &ListNode{
						value: 5,
						next:  nil,
					},
				},
			},
		},
	}

	fmt.Println(MthFromEnd(&ll, 0))
	fmt.Println(MthFromEnd(&ll, 1))
	fmt.Println(MthFromEnd(&ll, 2))
	fmt.Println(MthFromEnd(&ll, 3))
	fmt.Println(MthFromEnd(&ll, 4))

	//[1,3,6]
	fmt.Println("1->3->6")
	l2 := ListNode{
		value: 1,
		next: &ListNode{
			value: 3,
			next: &ListNode{
				value: 6,
				next:  nil,
			},
		},
	}

	fmt.Println(MthFromEnd(&l2, 1))
	fmt.Println(MthFromEnd(&l2, 0))

}

func TestStack_Min(t *testing.T) {
	var S = Stack{}
	S.Push(10)
	S.Push(3)
	S.Push(12)
	S.Push(2)
	fmt.Println(S.Min()) // prints 2
	S.Pop()
	fmt.Println(S.Min()) // prints 3
	S.Pop()
	fmt.Println(S.Min()) // prints 3

	fmt.Println("==============================")
	var S2 = Stack{}
	S2.Push(-2)
	S2.Push(0)
	S2.Push(-3)
	fmt.Println(S2.Min()) // prints -3
	S2.Pop()
	fmt.Println(S2.Min()) // prints -2
}

func TestRemoveRowsColumns(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 5, 10, 11},
	}

	newMatrix := RemoveRowsColumns(matrix, 5)
	fmt.Println(newMatrix)
	matrix2 := [][]int{
		{0, 1, 2, 3},
		{4, 5, 5, 7},
		{4, 2, 6, 7},
		{4, 2, 5, 7},
		{8, 1, 10, 11},
	}
	newMatrix2 := RemoveRowsColumns(matrix2, 5)
	fmt.Println(newMatrix2)

}
