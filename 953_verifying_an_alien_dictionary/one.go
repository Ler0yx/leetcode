package verifyingAnAlienDictionary

func isAlienSorted(words []string, order string) bool {
    
    inOrder := true
    
    //Iterates over the given Slice words
    for i:= 0; i < len(words) - 1; i++ {
        
        //Iterates over the Elements contained in the given Slice
        for d := 0; d < len(words[i]); d++ {
 
            //If there is one more Element (letter) then the one we are looking at right now
            if d < len(words[i+1]) {
                
                //Set x1 to the letter of the first word and x2 to the same letter of the second word
                x1 := words[i][d]
                x2 := words[i+1][d]
                
                //order is never longer than 26, so we set it to 27 to compare later
                o1 := 27
                o2 := 27
                
                //Iterates over the order string and compares the positions of the 2 letters x1 and x2
                for i := 0; i < len(order); i++ {
                    if x1 == order[i] {
                        o1 = i
                    }
                    if x2 == order[i] {
                        o2 = i
                    }
                }
                
                //If the letter of word one is first in order > jump to the next comparison
                //If the letter it is the same letter > compare the next letter
                //If the letter is second in order > set inOrder to false and return it
                if o1 != 27 && o2 != 27 {
                        if o1 == o2 {
                            continue
                        } else if o1 < o2 {
                            break
                        } else if o1 > o2 {
                            return !inOrder
                    }
                }
            } else {
                return !inOrder   
            }
        }
    }
    return inOrder
}