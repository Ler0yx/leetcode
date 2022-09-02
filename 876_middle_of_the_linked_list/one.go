package middleOfTheLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {

	solution := []*ListNode{}

	for i := 0; head != nil; i++ {
		solution = append(solution, head)
		head = head.Next
	}

	if len(solution)%2 == 0 {
		return solution[(len(solution)+1)/2]
	} else {
		return solution[len(solution)/2]
	}
}
