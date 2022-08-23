package main

func minCostClimbingStairs3(cost []int) int {

	costTotal := 0

	for i := 0; ; {
		switch {

		case i <= len(cost)-3 && cost[i]+cost[i+2] < cost[i+1]:
			costTotal += cost[i]
			i += 1
			if i <= len(cost)-4 && cost[i]+cost[i+2] > cost[i+1]+cost[i+3] {
				costTotal -= cost[i]
				costTotal += cost[i+1]
				i += 1
			}

		case i <= len(cost)-3 && cost[i]+cost[i+2] > cost[i+1]:
			costTotal += cost[i+1]
			i += 2
			if i <= len(cost)-4 && cost[i]+cost[i+2] < cost[i+1]+cost[i+3] {
				costTotal -= cost[i+1]
				costTotal += cost[i]
				i -= 1
			}

		case i <= len(cost)-2 && cost[i] < cost[i+1]:
			costTotal += cost[i]
			i += 1

		case i <= len(cost)-2 && cost[i] > cost[i+1]:
			costTotal += cost[i+1]
			i += 2

		case i > len(cost)-2:
			return costTotal

		}
	}
	return costTotal
}
