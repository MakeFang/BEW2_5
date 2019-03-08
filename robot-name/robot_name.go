package robotname

import (
    "math/rand"
    // "errors"
)

var seenName = map[string]int{}

type Robot struct {
  // Seen map[string]int
  robotName string
}

func (r *Robot) Rand () (string) {
  var result string
  for i := 0; i < 2; i++ {
    result = result + string(65 + rand.Intn(26))
  }
  for i := 0; i < 3; i++ {
    result = result + string(48 + rand.Intn(10))
  }
  return result
}

func (r *Robot) Name () (string, error) {
  if r.robotName != "" {
    return r.robotName, nil
  }
  result := r.Rand()
  for seenName[result] > 0 {
    result = r.Rand()
  }
  r.robotName = result
  seenName[result] += 1
  return r.robotName, nil
}

func (r *Robot) Reset () (string, error) {
  // seenName[r.robotName] -= 1
  result := r.Rand()
  for seenName[result] > 0 {
    result = r.Rand()
  }
  seenName[result] += 1
  r.robotName = result
  return r.robotName, nil
}
