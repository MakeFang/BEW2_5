package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func FrequencyCon(s string, c chan FreqMap) {
  m := FreqMap{}
	for _, r := range s {
    // m := <- c
		m[r]++
    // c <- m
	}
  c <- m
}

func ConcurrentFrequency(s []string) FreqMap {
  m := FreqMap{}
  lenS := len(s)
  var c chan FreqMap = make(chan FreqMap)
	for _, i := range s {
    // fmt.Println(i)
    go FrequencyCon(i, c)
  }
  // res := <- c
  // return res
  for i := 0; i < lenS; i++{
    msg := <- c
    // fmt.Println(msg)
    for k, v := range msg {
      m[k] += v
    }
    // fmt.Println("result: ", m)
  }
  close(c)
  return m
}
