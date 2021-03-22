// Package utils provides functions frequently used in go.
package sliceutils

import (
  "bytes"
)

func EqualIntSlice(a, b []int) bool {
  if len(a) != len(b) {
    return false
  }
  for i, v := range a {
    if v != b[i] {
      return false
    }
  }
  return true
}

func EqualByteMatrix(a, b [][]byte) bool {
  if len(a) != len(b) {
    return false
  }
  for i, v := range a {
    if !bytes.Equal(v, b[i]) {
      return false
    }
  }
  return true
}
