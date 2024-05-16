package leapyear

func isLeap(year int) bool {
    leap := false
 
    if year%4 == 0 {
        if year%100 == 0 {
            if year%400 == 0 {
                leap = true
            }
        }
        
        leap = true
    }
    
    return leap
}
