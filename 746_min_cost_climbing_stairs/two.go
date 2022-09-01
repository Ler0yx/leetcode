package minMostClimbingStairs

func minCostClimbingStairs2(cost []int) int {

	costTotal := 0

	if len(cost) == 2 {
		if cost[0] < cost[1] {
			return cost[0]
		} else {
			return cost[1]
		}
	}

	//Decision to start at 0.
	if cost[0]+cost[2] < cost[1] {

		for i := 0; ; {
			if i <= len(cost)-3 && cost[i]+cost[i+2] < cost[i+1] {
				costTotal += cost[i]
				i += 1
				if i <= len(cost)-4 && cost[i]+cost[i+2] > cost[i+1]+cost[i+3] {
					costTotal -= cost[i]
					costTotal += cost[i+1]
					i += 2
				}
			} else if i <= len(cost)-3 && cost[i]+cost[i+2] > cost[i+1] {
				costTotal += cost[i+1]
				i += 2
			} else {
				if cost[i] < cost[i+1] {
					costTotal += cost[i]
					return costTotal
				} else {
					costTotal += cost[i+1]
					return costTotal
				}
			}
		}
		//Decision to start at 1.
	} else {
		for i := 1; i < len(cost); {
			if i <= len(cost)-3 && cost[i]+cost[i+2] < cost[i+1] {
				costTotal += cost[i]
				i += 1
				if i <= len(cost)-4 && cost[i]+cost[i+2] > cost[i+1]+cost[i+3] {
					costTotal -= cost[i]
					costTotal += cost[i+1]
					i += 2
				}
			} else if i <= len(cost)-3 && cost[i]+cost[i+2] > cost[i+1] {
				costTotal += cost[i+1]
				i += 2
			} else {
				if cost[i] < cost[i+1] {
					costTotal += cost[i]
					return costTotal
				} else {
					costTotal += cost[i+1]
					return costTotal
				}
			}

		}
	}
	return costTotal
}
