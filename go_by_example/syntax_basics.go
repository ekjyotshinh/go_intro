package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// Constants are declared using the `const` keyword. They are compile-time constructs,
// meaning their value must be known at compile time. They can be character, string,
// boolean, or numeric values.
const s string = "constant"

func main() {
 	basicTypesExample()
	arrayExample()
	slicesExample()
	mapExample()
}

func basicTypesExample() {
	// A classic starting point. `fmt.Println` is used for basic output.
	log.Println("Demonstrating basic Go syntax and constructs.")

	// --- STRINGS & BASIC ARITHMETIC ---
	// Go supports standard string and arithmetic operations.
	log.Println("--- Strings and Arithmetic ---")
	fmt.Println("String Concatenation:", "go"+"lang")
	fmt.Println("Integer Addition (1+1):", 1+1)
	fmt.Println("Integer Subtraction (7-3):", 7-3)
	fmt.Println("Integer Multiplication (2*3):", 2*3)
	fmt.Println("Integer Division (8/4):", 8/4)
	fmt.Println("Integer Modulus (7%3):", 7%3)
	// Floating-point division behaves as expected.
	fmt.Println("Floating-Point Division (7.0/3.0):", 7.0/3.0)

	// --- BOOLEAN LOGIC & COMPARISONS ---
	// Standard boolean operators and comparison operators are available.
	log.Println("--- Boolean Logic and Comparisons ---")
	fmt.Println("Equality (3 == 3):", 3 == 3)
	fmt.Println("Inequality (3 != 4):", 3 != 4)
	fmt.Println("Less Than (3 < 4):", 3 < 4)
	fmt.Println("Less Than or Equal To (3 <= 3):", 3 <= 3)
	fmt.Println("Greater Than (4 > 3):", 4 > 3)
	fmt.Println("Greater Than or Equal To (4 >= 4):", 4 >= 4)
	fmt.Println("Logical AND (true && false):", true && false)
	fmt.Println("Logical OR (true || false):", true || false)
	fmt.Println("Logical NOT (!true):", !true)

	// --- VARIABLES ---
	// `var` declares 1 or more variables. You can declare a variable and initialize it later,
	// or do both at the same time.
	log.Println("--- Variable Declarations ---")
	var a = "initial"
	fmt.Println("Initialized variable 'a':", a)

	// Multiple variables can be declared at once.
	var b, c int = 1, 2
	fmt.Println("Multiple declarations 'b', 'c':", b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println("Type-inferred variable 'd':", d)

	// Variables declared without a corresponding initialization are zero-valued.
	// For example, the zero value for an `int` is `0`.
	var e int
	fmt.Println("Zero-valued integer 'e':", e)

	// The `:=` syntax is shorthand for declaring and initializing a variable.
	// It's only available within functions.
	f := "apple"
	fmt.Println("Shorthand declaration 'f':", f)

	// --- CONSTANTS (CONTINUED) ---
	// A deeper look at constants and their properties.
	log.Println("--- Constants ---")
	fmt.Println("Using a package-level constant 's':", s)

	// A numeric constant has no type until it's given one, such as by an
	// explicit conversion.
	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision.
	const D = 3e20 / n
	fmt.Println("Constant expression 'D':", D)

	// A numeric constant can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call.
	// Here, we perform a conversion to int64.
	fmt.Println("Constant 'D' converted to int64:", int64(D))

	// Bitwise operations can be used with constants.
	// `1 << 100` creates a very large number.
	const (
		Big   = 1 << 100
		Small = Big >> 99 // `>>` is the right bit shift operator.
	)
	// `Small` becomes `(1 << 100) >> 99`, which is `1 << 1`, or `2`.
	fmt.Println("Bitwise constant 'Small':", Small)

	// A constant can be used as a function argument. `math.Sin` expects a `float64`.
	fmt.Println("Using constant 'n' in math.Sin:", math.Sin(n))

	// --- LOOPS ---
	// Go has only one looping construct: the `for` loop.
	log.Println("--- Loops ---")

	// 1. The "while" loop style:
	// A `for` loop can act like a `while` loop from other languages.
	fmt.Println("'while' style loop:")
	i := 1
	for i <= 3 {
		fmt.Println("  - iteration:", i)
		i = i + 1
	}

	// 2. The classic `for` loop:
	// Initialization; condition; post-statement.
	fmt.Println("Classic 'for' loop:")
	for j := 0; j < 3; j++ {
		fmt.Println("  - iteration:", j)
	}

	// 3. The `for...range` loop:
	// This form iterates over a variety of data structures.
	// Here, `range 3` is a shorthand that provides values 0, 1, 2.
	fmt.Println("'for...range' loop:")
	for k := range 3 {
		fmt.Println("  - range:", k)
	}

	// 4. The infinite loop:
	// A `for` loop without a condition will loop repeatedly until you
	// `break` out of it or `return` from the enclosing function.
	fmt.Println("Infinite loop with 'break':")
	for {
		fmt.Println("  - in the loop")
		break // immediately exits the loop
	}

	// 5. Loop with `continue`:
	// `continue` skips to the next iteration of the loop.
	fmt.Println("Loop with 'continue' to skip even numbers:")
	for n := range 6 {
		if n%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("  - odd number:", n)
	}

    // --- IF/ELSE STATEMENTS ---
    // Go's `if` statements are straightforward. Parentheses are not required around the condition,
    // but braces are mandatory.
    log.Println("--- If/Else Statements ---")

    // 1. Simple `if` statement:
    if 7%2 == 0 {
        log.Println("7 is even")
    } else {
        log.Println("7 is odd")
    }

    // 2. `if` with a short statement:
    // You can include a short statement to execute before the condition.
    if num := 9; num < 0 {
        log.Println(num, "is negative")
    } else if num < 10 {
        log.Println(num, "has one digit")
    } else {
        log.Println(num, "has multiple digits")
    }   
    // Note : There is no ternary operator in Go. All conditional logic must be expressed with `if`, `else if`, and `else`.

    // --- SWITCH STATEMENTS ---
    // Go's `switch` statements are powerful and can be used in various ways.
    log.Println("--- Switch Statements ---")

    // 1. Simple `switch` statement:
    switch time.Now().Weekday() {
    case 1:
        log.Println("It's Monday!")
    case 2:
        log.Println("It's Tuesday!")
    default:
        log.Println("It's some other day.")
    }

    // 2. `switch` with multiple expressions in a case:
    day := time.Now().Weekday()
    switch day {
    case 0, 6:
        log.Println("It's the weekend!")
    default:
        log.Println("It's a weekday.")
    }

    // 3. `switch` without an expression:
    // This acts like a series of `if...else if` statements.
    t := time.Now()
    switch {
    case t.Hour() < 12:
        log.Println("Good morning!")
    case t.Hour() < 18:
        log.Println("Good afternoon!")
    default:
        log.Println("Good evening!")
    }
    
    // 4. Type switch:
    var k interface{} = 1
    switch v := k.(type) {
    case int:
        log.Println("Type is int:", v)
    case string:
        log.Println("Type is string:", v)
    default:
        log.Println("Type is unknown:", v)
    }

}
func arrayExample() {
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

func slicesExample() {
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
}

func mapExample() {
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
	value, exists := m["banana"]	// second value indicates if key is present
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

// functions with return types
// single return value
func add(a int, b int) int { // return types required
	return a + b
}

// multiple return values
func swap(x, y string) (string, string) { // add return types -- as many as required
	return y, x
}