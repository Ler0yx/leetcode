package studentAttendanceRecordI

func checkRecord(s string) bool {
    var absent int

    if len(s) == 1 {
        return true
    }
    for i := 0; i < len(s); i++ {
        switch {
            case s[i] == 65:
            absent++
            
            case i < len(s)-2 && (s[i] == 76 && s[i+1] == 76 && s[i+2] == 76):
            return false
        }
        if absent >= 2 {
            return false
        }
    }
    return true
}