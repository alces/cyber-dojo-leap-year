package leapyear

func isLeap(year int) (leap bool) { 
    if year%4 == 0 {
        leap = year%100 != 0 || year%400 == 0
    }
    
    return leap
}
