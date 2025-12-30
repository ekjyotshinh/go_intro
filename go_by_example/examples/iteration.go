package examples

import "fmt"

//range on arrays and slices provides both the index and value for each entry.
func iterateOverArray(){
	arr := [5]int{1,2,3,4,5}
	sum := 0
	for _,num :=range arr {	//ignoring index using _
		sum += num
	}
	fmt.Println("Sum is ", sum)
}

func iterateOverSlice(){
	slice := []string{"apple", "banana", "cherry"}
	for i, fruit := range slice {
		fmt.Printf("Index: %d, Fruit: %s\n", i, fruit)
	}
}

//range on maps provides key and value for each entry.
func iterateOverMap(){
	m := map[string]int{"a":1, "b":2, "c":3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

//range on strings iterates over Unicode code points.
func iterateOverString(){
	str := "hello"
	for i, ch := range str {
		fmt.Printf("Index: %d, Character: %c\n", i, ch)
	}
}

func DemoIteration(){
	iterateOverArray()
	iterateOverSlice()
	iterateOverMap()
	iterateOverString()
}