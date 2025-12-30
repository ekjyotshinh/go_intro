package main

import (
	"fmt"
	"go_by_example/examples"
)

func main() {
	// Uncomment the function calls below to see examples of various Go syntax basics.
	examples.BasicTypesExample()
	examples.ArrayExample()
	examples.SlicesExample()
	examples.MapExample()
	examples.UseVariadicFunc()
	fistInt := examples.AnonymousFuncExample()
	fmt.Println("Anonymous function calls:", fistInt(), fistInt(), fistInt())

	secondInt := examples.AnonymousFuncExample()
	fmt.Println("Another anonymous function calls:", secondInt(), secondInt())
}
