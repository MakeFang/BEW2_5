package grains

import "errors"

func Square (index int) (value uint64, err error) {
  if index > 64 || index <= 0 {
    return 0, errors.New("ERR")
  } else {
    indexNew := uint64(index)
    var result uint64
    result = 1 << (indexNew-1)
    return result, nil
  }

}

func Total () (total uint64) {
  return ^uint64(0)
}
