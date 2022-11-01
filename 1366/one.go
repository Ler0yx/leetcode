package rankTeamByVotes

func rankTeams(votes []string) string {
    positions := make(map[byte]map[int]int)
    var solution string
    var countVotes = len(positions[65])
    
    for i := 65; i <= 90; i++ {
        positions[byte(i)] = make(map[int]int)
    } 

    for _, vote := range votes {
        for i := 0; i < len(vote); i++ {
            positions[vote[i]][i] += 1
        }
    }
    fmt.Println(len(positions))

    for i := 65;; i++ {
        tmp := byte(i)
        for j := 0; j < len(positions[65]); j++ {
            if positions[tmp][j] > positions[tmp+1][j] {
                solution += string(tmp)
                positions[tmp][j] = 0
            } else if positions[tmp][j] < positions[tmp+1][j] {
                solution += string(tmp+1)
                positions[tmp+1][j] = 0
            }
            if i == countVotes {
                solution += string(tmp)
            }
        }
        if len(votes[0]) > len(solution) && i == 90 {
            i = 65
        } else if len(votes[0]) == len(solution) {
            break
        }
    }
    
    return solution
}