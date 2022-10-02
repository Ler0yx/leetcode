package keepMultiplyingFoundValuesByTwo

func findFinalValue(nums []int, original int) int {

loop:
	for i := 0; ; i++ {
		if nums[i] == original {
			original *= 2
			goto loop
		}
		if i == len(nums)-1 {
			return original
		}
	}
}
