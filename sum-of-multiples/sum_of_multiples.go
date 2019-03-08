package summultiples

func SumMultiples (limit int, divisors []int) int {
  if len(divisors) == 0 {
    return 0
  }
  results := map[int]int{}
  result := 0
  // fmt.Println("entered")
  for _, divisor := range divisors {
    if divisor == 0 {
      continue
    }
    // fmt.Println(divisor)
    curPow := 0
    curNum := curPow * divisor
    for curNum < limit{
      if results[curNum] == 0 {
        results[curNum] += 1
        result += curNum
      }
      curPow += 1
      curNum = curPow * divisor
    }
  }
  return result
}
