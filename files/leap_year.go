package leapyear

func isLeap(year int) bool {
    if year%4 == 0 {
        if year%100 == 0 {
            return false
        }
        
        return true
    }
    
    return false
}
