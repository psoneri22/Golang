package main

// Define an interface named Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct named Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement methods to satisfy the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Define a struct named Circle
type Circle struct {
	Radius float64
}

// Implement methods to satisfy the Shape interface for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// func main() {
// 	// Create instances of Rectangle and Circle
// 	rectangle := Rectangle{Width: 3, Height: 4}
// 	circle := Circle{Radius: 5}

// 	// Call methods on instances of Rectangle and Circle
// 	fmt.Println("Rectangle Area:", rectangle.Area())
// 	fmt.Println("Rectangle Perimeter:", rectangle.Perimeter())

// 	fmt.Println("Circle Area:", circle.Area())
// 	fmt.Println("Circle Perimeter:", circle.Perimeter())

// 	// Assign Rectangle and Circle instances to Shape interface variables
// 	var shape Shape
// 	shape = rectangle
// 	fmt.Println("Shape Area:", shape.Area())

// 	shape = circle
// 	fmt.Println("Shape Area:", shape.Area())
// }
