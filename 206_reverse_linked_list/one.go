package reverseLinkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    list := make([]int, 0)
    var ctr int
    for i := head; i != nil; i = i.Next {
        list = append(list, i.Val)
    }
    ctr = len(list)-1
    for j := head; j != nil; j = j.Next {
        j.Val = list[ctr]
        ctr--
    }
    return head
}