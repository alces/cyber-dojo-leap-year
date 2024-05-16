package leapyear

func isLeap(year int) bool {
    if year%4 == 0 {
        return true
    }
    
    return false
}
