package shortestDistanceToTargetStringInACircularArray

func closetTarget(words []string, target string, startIndex int) int {

	var (
		dst int
		ctr bool
		n   int = len(words)
	)

	for i, j := startIndex, startIndex; ; i, j, ctr, dst = (i+1)%n, (j-1)%n, true, dst+1 {
		switch {
		case words[i] == target || words[j] == target:
			return dst

		case (ctr && i == startIndex) || (n == 1 && words[0] != target):
			return -1

		case i == len(words)-1:
			i = -1

		case j == 0:
			j = n
		}
	}
	return 0
}
