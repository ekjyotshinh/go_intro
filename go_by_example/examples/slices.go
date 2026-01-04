package examples

import (
	"fmt"
	"log"
	"slices"
)

func SlicesExample() {
	// --- SLICES ---
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	log.Println("--- Slices ---")
	var s []string;
	fmt.Println("Empty slice 's':", s) // len(s) == 0 and s == nil
	
	// 1. Creating a slice with `make`:
	s = make([]string, 3)
	fmt.Println("Slice 's' created with make:", s)
	fmt.Println("Length of slice 's':", len(s))
	fmt.Println("Capacity of slice 's':", cap(s))
	
	// 2. Initializing a slice with values:
	s = []string{"apple", "banana", "cherry"}
	fmt.Println("Slice 's' initialized with values:", s)
	
	// 3. Appending to a slice:
	s = append(s, "date")
	fmt.Println("Slice 's' after appending 'date':", s)
	
	// 4. Slicing a slice:
	subSlice := s[1:3]
	fmt.Println("Sub-slice of 's' from index 1 to 3:", subSlice)
	
	// 5. Iterating over a slice using `for...range`:
	fmt.Println("Iterating over slice 's':")
	for i, v := range s {
		fmt.Printf("  - Index %d: Value %s\n", i, v)
	}

	// 6. Copying a slice:

	copiedSlice := make([]string, len(s))
	copy(copiedSlice, s)
	fmt.Println("Copied slice 'copiedSlice':", copiedSlice)

	// 7. Multi-dimensional slice:

	twoDSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D Slice 'twoDSlice':", twoDSlice)

	//soritng a slice
	strs := []string{"c", "a", "b"}
    slices.Sort(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    slices.Sort(ints)
    fmt.Println("Ints:   ", ints)

    s2 := slices.IsSorted(ints)
    fmt.Println("Sorted: ", s2)
}
