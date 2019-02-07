package isogram

import "strings"

func IsIsogram(input string) bool {
    var input_lower string = strings.ToLower(input)
    counts := make(map[rune]int)
    for _, cur_rune := range input_lower {
        if string(cur_rune) == "-" || string(cur_rune) == " " {
            continue
        }
        _, ok := counts[cur_rune]
        if ok {
            return false
        } else {
            counts[cur_rune] += 1
        }
    }
    return true
}
