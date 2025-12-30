package examples

import "fmt"

// recursive function calls itself with a modified argument until it reaches a base case.
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func DemoRecursion() {
    fmt.Println(fact(7))

	//Anonymous functions can also be recursive, but this requires explicitly declaring a variable with var to store the function before it’s defined.
    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }

        return fib(n-1) + fib(n-2)
    }

    fmt.Println(fib(7))
}