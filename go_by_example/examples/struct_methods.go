package examples

import "fmt"

type rectangle struct {
	width, height int
}

// we could also use a value receiver here, but using a pointer receiver is more efficient if the struct is large or if we want to modify the struct's fields.
func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) perimeter() int {
	return 2 * (r.width + r.height)
}

func DemoStructMethods() {
	rect := rectangle{width: 10, height: 5}
	area := rect.area()
	perimeter := rect.perimeter()
	fmt.Println("Area of rectangle:", area)
	fmt.Println("Perimeter of rectangle:", perimeter)
}