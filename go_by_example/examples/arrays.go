package examples

import (
	"fmt"
	"log"
)

func ArrayExample() {
	// --- ARRAYS ---
	// An array is a numbered sequence of elements of a specific length.
	log.Println("--- Arrays ---")

	// 1. Declaring and initializing an array:
	var a [5]int
	fmt.Println("Array 'a' with default values:", a)

	// 2. Initializing an array with values:
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array 'b' initialized with values:", b)

	// 3. Accessing and modifying array elements:
	b[0] = 10
	fmt.Println("Modified array 'b':", b)

	// 4. Getting the length of an array:
	fmt.Println("Length of array 'b':", len(b))

	// 5. Iterating over an array using `for...range`:
	fmt.Println("Iterating over array 'b':")
	for i, v := range b {
		fmt.Printf("  - Index %d: Value %d\n", i, v)
	}

	// 6. Array with inferred length:
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array with inferred length 'b':", b)

	//2d array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D Array 'twoD':", twoD)
}
