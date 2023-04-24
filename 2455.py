# 2455. Average Value of Even Numbers That Are Divisible by Three

class Solution:
    def averageValue(self, nums: List[int]) -> int:
        sum, ctr = 0, 0
        for i in range(len(nums)):
            if nums[i]%2 == 0 and nums[i]%3 == 0:
                sum += nums[i]
                ctr += 1
        if ctr == 0:
            return 0
        return int(sum/ctr)
