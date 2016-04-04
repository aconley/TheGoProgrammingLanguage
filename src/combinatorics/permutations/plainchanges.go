package permutations

import "sort"

// PlainChanges visits all the permutations
// of data.  Requires that the elements
// are distinct.  Data is modified by this process
func PlainChanges(data Interface) {

  n := data.Len()
  if n == 0 {
    return
  }
  if !sort.IsSorted(data) {
    sort.Sort(data)
  }

  // Visit the initial permutation
  if data.Visit() {
      return
  }

  // Algorithm P of Knuth Volume 4A 7.2.1.2
  c := make([]int, n)
  o := make([]int, n)
  for i := range(o) {
    o[i] = 1
  }

  // Main loop
  for {
    // Find the index j to change
    j := n - 1
    s := 0
    for q := c[j] + o[j];; {
      if q == j {
        if j == 1 {
          // No more valid perms
          return
        }
        s++
        o[j] = -o[j]
        j--
      } else if q < 0 {
        o[j] = -o[j]
        j--
      } else {
        // Visit, possibly terminate
        if data.Visit() {
          return
        }
        c[j] = q
      }
    }
  }
}