package luhn

import "strings"

func Valid(inputString string) bool {
    stripString := strings.Replace(inputString, " ", "", -1)
    rs := []rune(stripString)
    strLen := len(rs)
    if strLen <= 1 {
        return false
    }
    totalSum := 0
    counter := 0
    for i := strLen-1 ; i >= 0; i-- {
        if '0' <= rs[i] && rs[i] <= '9' {
            if counter % 2 == 0 {
                resInt := int(rs[i]) - '0'
                totalSum += resInt
            } else {
                resInt := (int(rs[i]) - '0')*2
                if resInt > 9 {
                    resInt -= 9
                }
                totalSum += resInt
            }
        } else {
            return false
        }
        counter += 1
    }
    if totalSum % 10 == 0 {
        return true
    } else {
        return false
    }
    return true
}
