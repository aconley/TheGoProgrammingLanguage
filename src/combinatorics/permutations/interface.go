package permutations

// Interface can be permuted by the routines in this package.
//  Types that satisfy this almost satisfy sort as well,
//  except that they don't have a comparison operator
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int

    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)

    // Visit is called on each permutation.  If it returns
    //  false, no more permutations are generate
    Visit() bool
}

// SortableInterface are both permutable and sortable;
//  that is, they satisfy both Interface and sort.Interface
type SortableInterface interface {
  Interface
  // Less reports whether the element with
  // index i should sort before the element with index j.
  Less(i, j int) bool
}