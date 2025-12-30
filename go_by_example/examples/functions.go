package examples

import "fmt"

// functions with return types
// single return value
func Add(a int, b int) int { // return types required
	return a + b
}

// multiple return values
func Swap(x, y string) (string, string) { // add return types -- as many as required
	return y, x
}

// variadic function can be called with any number of arguments
func Sum(nums ...int) int { // variadic function
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func UseVariadicFunc() {
	fmt.Println("Sum of 1, 2, 3:", Sum(1, 2, 3))		   // prints 6
	fmt.Println("Sum of 10, 20, 30, 40, 50:", Sum(10, 20, 30, 40, 50)) // prints 150

	nums := []int{5, 10, 15}
	fmt.Println("Sum of nums slice:", Sum(nums...)) // prints 30
}

// ananymous functions
// return a function that increments and returns an integer
// in this case the value of i in enclosed in the anonymous function's closure
// so each time the returned function is called, it has access to its own i variable
func AnonymousFuncExample() func() int {
	i := 0
	return func() int{
		i++
		return i
	}
}
