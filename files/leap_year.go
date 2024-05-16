package leapyear

func isLeap(year int) (leap bool) { 
    if year%4 == 0 {
        if year%100 == 0 {
           leap = year%400 == 0
        } else {
            leap = true
        }
    }
    
    return leap
}
