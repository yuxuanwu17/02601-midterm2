package main

import (
	"fmt"
	"strconv"
)

//======================Q1======================

func main() {
	node7 := TreeNode{
		left:  nil,
		right: nil,
		value: 7,
	}

	node4 := TreeNode{
		left:  nil,
		right: nil,
		value: 4,
	}

	node3 := TreeNode{
		left:  &node4,
		right: &node7,
		value: 3,
	}

	node0 := TreeNode{
		left:  nil,
		right: nil,
		value: 0,
	}

	node2 := TreeNode{
		left:  &node4,
		right: &node0,
		value: 2,
	}

	node8 := TreeNode{
		left:  &node2,
		right: &node7,
		value: 8,
	}

	node13 := TreeNode{
		left:  nil,
		right: nil,
		value: 13,
	}

	node12 := TreeNode{
		left:  &node13,
		right: nil,
		value: 12,
	}

	node1 := TreeNode{
		left:  nil,
		right: nil,
		value: 1,
	}

	node9 := TreeNode{
		left:  &node1,
		right: &node12,
		value: 9,
	}

	node5 := TreeNode{
		left:  &node4,
		right: nil,
		value: 5,
	}

	node6 := TreeNode{
		left:  &node5,
		right: &node9,
		value: 6,
	}

	node23 := TreeNode{
		left:  nil,
		right: nil,
		value: 23,
	}
	node15 := TreeNode{
		left:  nil,
		right: nil,
		value: 15,
	}

	node16 := TreeNode{
		left:  nil,
		right: nil,
		value: 16,
	}

	node19 := TreeNode{
		left:  &node15,
		right: nil,
		value: 19,
	}

	node20 := TreeNode{
		left:  &node19,
		right: &node23,
		value: 20,
	}

	node10 := TreeNode{
		left:  &node5,
		right: nil,
		value: 10,
	}

	node14 := TreeNode{
		left:  &node10,
		right: &node20,
		value: 14,
	}

	node17 := TreeNode{
		left:  &node16,
		right: &node23,
		value: 17,
	}
	node22 := TreeNode{
		left:  &node10,
		right: &node17,
		value: 22,
	}

	fmt.Println("==================Test Q1============================")
	fmt.Println(IsBSTOrdered(&node3))  // F
	fmt.Println(IsBSTOrdered(&node8))  // F
	fmt.Println(IsBSTOrdered(&node6))  // F
	fmt.Println(IsBSTOrdered(&node14)) // T
	fmt.Println(IsBSTOrdered(&node22)) // F
	fmt.Println("")

	fmt.Println("==================Test Q2============================")
	fmt.Println(LowestCommonAncestor(&node14, 14, 10)) // 14
	fmt.Println(LowestCommonAncestor(&node14, 19, 23)) // 20
	fmt.Println(LowestCommonAncestor(&node14, 4, 15))  // 14
	fmt.Println(LowestCommonAncestor(&node14, 15, 23)) // 20
	fmt.Println(LowestCommonAncestor(&node14, 10, 4))  // 10
	fmt.Println("")

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
	lnode0 := ListNode{
		next:  &lnode2_2,
		value: 0,
	}

	//fmt.Println("==================Test Q3============================")
	//fmt.Println(printLinkedList(&lnode1))                                  // 1 3 6
	//fmt.Println(printLinkedList(&lnode2_2_2_2))                            // 2 2 2 2 6
	//fmt.Println(printLinkedList(&lnode0))                                  // 0 2 2 6
	//fmt.Println(printLinkedList(MergeSortedLists(&lnode1, &lnode2_2_2_2))) // 1 2 2 2 2 3 6 6
	//fmt.Println(printLinkedList(MergeSortedLists(&lnode1, &lnode0)))       // 0 1 2 2 3 6 6
	//fmt.Println(printLinkedList(MergeSortedLists(&lnode2_2_2_2, &lnode0))) // 0 2 2 2 2 2 2 6 6
	//fmt.Println(printLinkedList(&lnode0))                                  // 0 2 2 6
	//fmt.Println("")

	fmt.Println("==================Test Q4============================")
	fmt.Println(MthFromEnd(&lnode1, 1))       // 3
	fmt.Println(MthFromEnd(&lnode1, 0))       // 6
	fmt.Println(MthFromEnd(&lnode2_2_2_2, 0)) // 6
	fmt.Println(MthFromEnd(&lnode2_2_2_2, 4)) // 2
	fmt.Println(MthFromEnd(&lnode0, 3))       // 0
	fmt.Println(MthFromEnd(&lnode2, 1))       // 2
	fmt.Println("")

	fmt.Println("==================Test Q5============================")
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
	S.Pop()
	fmt.Println(S.Min()) // prints 10
	fmt.Println("")

	fmt.Println("==================Test Q6============================")
	m1 := [][]int{
		{3, 3, 2, 3},
		{4, 4, 6, 7},
		{8, 9, 12, 11},
	}
	m2 := [][]int{
		{3, 0, 2, 2},
		{4, 4, 6, 7},
		{8, 9, 3, 11},
	}

	m3 := [][]int{
		{3, 3, 3, 3},
		{4, 4, 6, 7},
		{8, 9, 12, 11},
	}

	m4 := [][]int{
		{3, 0, 2, 19},
		{3, 4, 6, 7},
		{8, 9, 12, 11},
	}

	mm1 := RemoveRowsColumns(m1, 3)
	mm2 := RemoveRowsColumns(m2, 3)
	mm3 := RemoveRowsColumns(m3, 3)
	mm4 := RemoveRowsColumns(m4, 3)
	fmt.Println(mm1) // [[6] [12]]
	fmt.Println(mm2) // [[4] [7]]
	fmt.Println(mm3) // [[] []]
	fmt.Println(mm4) // [[9 12 11]]
}

func printLinkedList(l *ListNode) string {
	if l.next == nil {
		k := strconv.Itoa(l.value)
		return k
	} else {
		o := strconv.Itoa(l.value)
		s := printLinkedList(l.next)
		str := o + "->" + s
		return str
	}
}
