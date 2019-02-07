package hamming

import "errors"

// If the strings have different lengths, raise error
// Loop through the two strings. If the characters are different, increase
// distance by 1.
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("Input string have different lengths.")
    }
    var distance int
    for index, char := range a {
        if string(char) != string(b[index]){
            distance += 1
        }
    }
    return distance, nil
}
