package raindrops

import "strconv"

func Convert(x int) string {
    var result string
    if x % 3 == 0 {
        result += "Pling"
    }
    if x % 5 == 0 {
        result += "Plang"
    }
    if x % 7 == 0 {
        result += "Plong"
    }
    if result != "" {
        return result
    }
    return strconv.Itoa(x)
}
