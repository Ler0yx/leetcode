package palindromeLinkedList

func isPalindrome(head *ListNode) bool {

	valMap := make(map[int]int)
	turn := 0

	for i := 0; head != nil; i++ {
		valMap[i] = head.Val
		head = head.Next

		if i > 0 && valMap[i] == valMap[i-1] {
			turn = i
		}
	}

	if turn != len(valMap)/2 {
		return false
	}

	if len(valMap) == 1 || len(valMap) == 2 && valMap[0] == valMap[1] {
		return true
	}

	for i := turn; i < len(valMap); {
		for j := turn - 1; j >= 0; j-- {
			if valMap[i] != valMap[j] {
				return false
			}
			i++
		}
	}
	return true
}


