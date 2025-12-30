package examples

import (
	"fmt"
	"log"
)

func MapExample() {
	// --- MAPS ---
	// A map is an unordered collection of key-value pairs.
	log.Println("--- Maps ---")

	// 1. Creating a map:
	m := make(map[string]int) // map with string keys and int values
	fmt.Println("Empty map 'm':", m)
	// 2. Adding key-value pairs:
	m["apple"] = 1
	m["banana"] = 2
	fmt.Println("Map 'm' after adding key-value pairs:", m)

	// 3. Accessing values by key:
	appleCount := m["apple"]
	fmt.Println("Value for key 'apple':", appleCount)

	// 4. Deleting a key-value pair:
	delete(m, "banana")
	fmt.Println("Map 'm' after deleting key 'banana':", m)

	// 5. Checking if a key exists:
	value, exists := m["banana"] // second value indicates if key is present
	if exists {
		fmt.Println("Key 'banana' exists with value:", value)
	} else {
		fmt.Println("Key 'banana' does not exist. value is default to :", value)
	}

	// 6. Iterating over a map using `for...range`:
	fmt.Println("Iterating over map 'm':")
	for k, v := range m {
		fmt.Printf("  - Key: %s, Value: %d\n", k, v)
	}

	// 7. clearing a map:
	clear(m)
	fmt.Println("Map 'm' after clearing all key-value pairs:", m)
}
