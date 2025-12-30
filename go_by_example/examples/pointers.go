package examples

import (
	"fmt"
	"log"
)

// pointers hold the memory address of a value.
// Pointers allow you to pass references to values and records within your program,
// which is essential for building efficient and modifiable applications.

// `zeroval` demonstrates pass-by-value.
// It receives a *copy* of the integer `ival`.
// Any changes made to `ival` inside this function do NOT affect the original variable.
func zeroval(ival int) {
	log.Printf("Inside zeroval: Received value %d. Setting it to 0.\n", ival)
	ival = 0
}

// `zeroptr` demonstrates pass-by-reference using a pointer.
// It receives a pointer to an integer, `*int`.
// The `*` in the parameter type indicates it's a pointer.
func zeroptr(iptr *int) {
	log.Printf("Inside zeroptr: Received pointer to memory address %v.\n", iptr)
	// The `*` operator here is used for DEREFERENCING.
	// Dereferencing means accessing the value that the pointer `iptr` points to.
	// By setting `*iptr = 0`, we are changing the value at the specified memory address.
	*iptr = 0
}

// DemoPointers runs through a detailed demonstration of pointer concepts.
func DemoPointers() {

	i := 1
	fmt.Printf("Initial value of 'i': %d\n", i)

	// The `&` operator gives you the memory address of a variable.
	// The type of `&i` is `*int` (a pointer to an integer).
	fmt.Printf("Memory address of 'i' (&i): %v\n", &i)


	fmt.Printf("Value of 'i' before calling zeroval: %d\n", i)
	zeroval(i) // A copy of `i`'s value (1) is passed.
	fmt.Printf("Value of 'i' after calling zeroval: %d (Unchanged)\n", i)


	fmt.Printf("Value of 'i' before calling zeroptr: %d\n", i)
	zeroptr(&i) // The memory address of `i` is passed.
	fmt.Printf("Value of 'i' after calling zeroptr: %d (Changed)\n", i)

	// You can declare a variable that holds a pointer.
	var p *int
	fmt.Printf("Declared a nil pointer 'p' of type *int. Value: %v\n", p) // A zero-value pointer is nil.

	// Assign the memory address of `i` to the pointer `p`.
	p = &i
	fmt.Printf("Pointer 'p' now points to the address of 'i': %v\n", p)

	// You can see the value stored at a pointer's address by dereferencing it with `*`.
	fmt.Printf("Value at the address 'p' points to (*p): %d\n", *p)

	// If you change the value through the pointer, the original variable also changes.
	fmt.Println("Changing value through the pointer 'p' to 42 (*p = 42).")
	*p = 42
	fmt.Printf("New value of 'i': %d\n", i)
	fmt.Printf("New value via pointer (*p): %d\n", *p)


	type person struct {
		name string
		age  int
	}

	// use pointers to structs to avoid copying large data structures.
	p1 := person{name: "Alice", age: 30}
	fmt.Printf("Original struct p1: %+v\n", p1)

	// Create a pointer to the struct instance.
	p1_ptr := &p1
	fmt.Printf("Pointer to p1: %v\n", p1_ptr)

	// Go provides a convenient shortcut for accessing struct fields through a pointer.
	// Instead of writing `(*p1_ptr).name`, you can just write `p1_ptr.name`.
	fmt.Printf("Accessing name via pointer (p1_ptr.name): %s\n", p1_ptr.name)

	// Modifying the struct's field through the pointer.
	fmt.Println("Modifying age via pointer (p1_ptr.age = 31).")
	p1_ptr.age = 31
	fmt.Printf("Modified struct p1: %+v (Original struct is changed)\n", p1)
}
