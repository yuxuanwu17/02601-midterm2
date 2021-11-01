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

func TestMergeSortedLists(t *testing.T) {
	l1 := ListNode{
		value: 1,
		next: &ListNode{
			value: 2,
			next: &ListNode{
				value: 4,
				next:  nil,
			},
		},
	}

	l2 := ListNode{
		value: 1,
		next: &ListNode{
			value: 3,
			next: &ListNode{
				value: 4,
				next:  nil,
			},
		},
	}

	mergedList := MergeSortedLists(&l1, &l2)
	for {
		if mergedList.next == nil {
			break
		}
		fmt.Println(mergedList.value)
		mergedList = mergedList.next
	}

	fmt.Println("==================")

	l12 := ListNode{}
	l22 := ListNode{}
	mergedList2 := MergeSortedLists(&l12, &l22)
	//fmt.Println(mergedList2)
	for {
		if mergedList2.next == nil {
			break
		}
		fmt.Println(mergedList2.value)
		mergedList2 = mergedList2.next
	}
	fmt.Println("==================")

	l13 := ListNode{value: 0, next: nil}
	mergedList3 := MergeSortedLists(&l12, &l13)
	fmt.Println(mergedList3.value)
}
