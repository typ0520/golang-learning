package main

import (
	"fmt"
)

func main() {
	// 2. 两数相加
	// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

	// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

	// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	// 示例：

	// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	// 输出：7 -> 0 -> 8
	// 原因：342 + 465 = 807

	// 来源：力扣（LeetCode）
	// 链接：https://leetcode-cn.com/problems/add-two-numbers
	// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

	// 99910,009,9989999+9999=10,009,998

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	// l1 := &ListNode{
	// 	Val: 9,
	// 	Next: &ListNode{
	// 		Val: 9,
	// 		Next: &ListNode{
	// 			Val: 9,
	// 			Next: &ListNode{
	// 				Val:  9,
	// 				Next: nil,
	// 			},
	// 		},
	// 	},
	// }
	// l2 := &ListNode{
	// 	Val:  9,
	// 	Next: nil,
	// }

	l3 := addTwoNumbers(l1, l2)

	var str string
	for l3 != nil {
		str = fmt.Sprintf("%d%s", l3.Val, str)
		l3 = l3.Next
	}
	fmt.Printf("%s", str)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for {
		if l1 != nil {
			current.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current.Val += l2.Val
			l2 = l2.Next
		}

		if current.Val > 9 {
			current.Val -= 10
			current.Next = &ListNode{
				Val: 1,
			}
		}
		if l1 == nil && l2 == nil {
			break
		}
		if current.Next == nil {
			current.Next = &ListNode{}
		}
		current = current.Next
	}
	return head
}
