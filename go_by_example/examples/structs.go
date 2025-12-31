package examples

import "fmt"
type person struct{
	name string
	age  int
}

// returns a pointer to the newly created person struct
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 10
	return &p
}

func DemoStructs() {
	fmt.Println("Structs Example")

	// Create a new person using the newPerson function
	p1 := newPerson("Alice")
	fmt.Println("New Person:", p1.name, "Age:", p1.age)

	// Create a person struct using a struct literal
	p2 := person{name: "Bob", age: 25}
	fmt.Println("Person Literal:", p2.name, "Age:", p2.age)

	// Access and modify struct fields
	p2.age += 1
	fmt.Println("After Birthday:", p2.name, "Age:", p2.age)

	
    fmt.Println(person{"Bob", 20})

    fmt.Println(person{name: "Alice", age: 30})

    fmt.Println(person{name: "Fred"})

    fmt.Println(&person{name: "Ann", age: 40})


    fmt.Println(newPerson("Jon"))

	// Anonymous struct -- useful for one-off data structures
	dog := struct {
        name   string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)
}