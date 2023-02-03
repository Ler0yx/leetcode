package reverseLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var vals []int
	for i := head; i != nil; i = i.Next {
		vals = append(vals, i.Val)
	}
	for j := head; j != nil; j = j.Next {
		j.Val = vals[len(vals)-1]
		vals = vals[:len(vals)-1]
	}
	return head
}
