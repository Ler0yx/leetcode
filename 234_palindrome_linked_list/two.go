package palindromeLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome2(head *ListNode) bool {

	values := []int{}

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	if len(values)%2 == 0 {
		for i, j := len(values)/2, len(values)/2-1; i < len(values); i++ {
			if values[i] != values[j] {
				return false

			}
			j--
		}
	} else {
		for i, j := len(values)/2+1, len(values)/2-1; i < len(values); i++ {
			if values[i] != values[j] {
				return false
			}
			j--
		}
	}
	return true
}
