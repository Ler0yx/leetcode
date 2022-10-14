package countItemsMatchingARule

func countMatches(items [][]string, ruleKey string, ruleValue string) int {

	var counter int

	for _, item := range items {
		if ruleKey == "type" && ruleValue == item[0] {
			counter += 1
			continue
		}
		if ruleKey == "color" && ruleValue == item[1] {
			counter += 1
			continue
		}
		if ruleKey == "name" && ruleValue == item[2] {
			counter += 1
			continue
		}
	}
	return counter
}
