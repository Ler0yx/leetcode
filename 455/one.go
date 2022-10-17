package assignCookies

import "fmt"

func findContentChildren(g []int, s []int) int {

	gMap := make(map[int]int)
	childs := make(map[int][]int)
	sMap := make(map[int]int)
	var distChilds []int
	counter := 0
	remaining := 0

	for i := 0; i < len(g); i++ {
		gMap[g[i]] += 1
		if _, ok := childs[g[i]]; !ok {
			childs[g[i]] = append(childs[g[i]], g[i])
			childs[0] = append(childs[0], g[i])
		}
	}
	for j := 0; j < len(s); j++ {
		sMap[s[j]] += 1
	}

	for k := 0; k < len(distChilds); k++ {
		fmt.Println(remaining)
		if gMap[distChilds[k]]+remaining <= sMap[distChilds[k]] {
			counter += gMap[distChilds[k]]
			remaining = 0
		} else if gMap[distChilds[k]]+remaining > sMap[distChilds[k]] {
			counter += sMap[distChilds[k]]
			remaining += gMap[distChilds[k]] - sMap[distChilds[k]]
		}
	}
	return counter
}
