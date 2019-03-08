package lsproduct

import "errors"

func LargestSeriesProduct (digits string, span int) (p int, err error) {
  size := len(digits)
  if span > size || span < 0 {
    return 0, errors.New("BAD")
  }
  maxNum := 1
  for _, start := range digits[:span] {
    maxNum = maxNum * (int(start) - '0')
  }
  for i:= 0; i < size - span + 1; i++ {
    product := 1
    // fmt.Println(digits[i:i+span])
    for _, comp := range digits[i:i+span] {
        if comp > '9' || comp < '0' {
          return 0, errors.New("BAD")
        }
      product = product * (int(comp) - '0')
    }
    if product > maxNum {
      maxNum = product
    }
  }
  return maxNum, nil
}
