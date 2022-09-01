package minMostClimbingStairs

func minCostClimbingStairs(cost []int) int {

	costTotal := 0

	if cost[0] > cost[1] {
		costTotal := cost[1]
		for i := 1; i < len(cost); i++ {

			//If there is no third number to consider for the most efficient path, just compare the next two numbers.
			if i >= len(cost)-3 {
				if cost[i+1] < cost[i+2] {
					costTotal += cost[i+1]
					return costTotal
				} else if cost[i+2] < cost[i+1] {
					costTotal += cost[i+2]
					return costTotal
				}
			} else if i > len(cost)-3 {
				return costTotal
			}

			//If there is a third number after the two we are comparing, we have to consider it into our evaluation to find the most efficient path.
			if (cost[i+1] + cost[i+3]) < cost[i+2] {
				costTotal += cost[i+1]
				i += 1
			} else {
				costTotal += cost[i+2]
				i += 2
			}
		}

		//If the first value is the smaller one, we have to consider the third number after that to find the most efficient path.
	} else if cost[1] > cost[0] {
		if (cost[0] + cost[2]) < cost[1] {
			costTotal := cost[0]
			for i := 0; i < len(cost); i++ {

				//If there is no third number to consider for the most efficient path, just compare the next two numbers.
				if i >= len(cost)-3 {
					if cost[i+1] < cost[i+2] {
						costTotal += cost[i+1]
						return costTotal
					} else if cost[i+2] < cost[i+1] {
						costTotal += cost[i+2]
						return costTotal
					}
				} else if i > len(cost)-3 {
					return costTotal
				}

				if cost[i+1] < cost[i+2] {
					if (cost[i+1] + cost[i+3]) < cost[i+2] {
						costTotal += cost[i+1]
						i += 1
					}
				} else if cost[i+2] < cost[i+1] {
					costTotal += cost[i+2]
					i += 2
				}
			}
		} else {
			costTotal := cost[1]
			for i := 1; i < len(cost); i++ {

				//If there is no third number to consider for the most efficient path, just compare the next two numbers.
				if i == len(cost)-3 {
					if cost[i+1] < cost[i+2] {
						costTotal += cost[i+1]
						return costTotal
					} else if cost[i+2] < cost[i+1] {
						costTotal += cost[i+2]
						return costTotal
					}
				} else if i > len(cost)-3 {
					return costTotal
				}

				if cost[i+1] < cost[i+2] {
					if (cost[i+1] + cost[i+3]) < cost[i+2] {
						costTotal += cost[i+1]
						i += 1
					}
				} else if cost[i+2] < cost[i+1] {
					costTotal += cost[i+2]
					i += 2
				}
			}
		}
	}

	return costTotal
}
