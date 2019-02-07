package isogram

import "strings"

func IsIsogram(input string) bool {
    var inputLower string = strings.ToLower(input)
    counts := make(map[rune]int)
    for _, cur_rune := range inputLower {
        if string(curRune) == "-" || string(curRune) == " " {
            continue
        }
        _, ok := counts[curRune]
        if ok {
            return false
        }
        counts[curRune] += 1
    }
    return true
}
