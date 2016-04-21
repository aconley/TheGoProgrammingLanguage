package permutations

// Heap visits all the permutations of data
//  using Heap's scheme.
// This is Algorithm G of Knuth Volum 4A 7.2.1.2
//  using the permutation of 7.2.1.2.(27)
func Heap(data Interface) {
  n := data.Len()
  if n == 0 {
    return
  }

  // Visit the initial permutation
  if !data.Visit() {
    return
  }

  // Setup
  c := make([]int, n + 1)

  for {
    k := 1
    for c[k] == k {
      c[k] = 0
      k++
    }
    if k == n {
      return
    }
    c[k]++
    if isEven(k) {
      data.Swap(0, k)
    } else {
      data.Swap(k, c[k] - 1)
    }

    if !data.Visit() {
      return
    }
  }
}

func isEven(x int) bool {
  return x&1 == 0;
}