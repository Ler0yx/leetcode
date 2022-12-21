package numberOfLinesToWriteString

func numberOfLines(widths []int, s string) []int {
    var lines int = 1
    var lastWidth int

    for i := 0; i < len(s); i++ {
        if lastWidth + widths[int(s[i])-97] > 100 {
            lastWidth = 0
            lines++
            i--
            continue
        }
        lastWidth += widths[int(s[i])-97]
    }   
    return []int{lines, lastWidth}
}