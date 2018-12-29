package main

import "fmt"

/*
 * structs are user defined types, in this example, 
 * we created a type called 'Vertex' within 2 variables
 */
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{17, 23} //here we are assigning 'v' with our new type values
	fmt.Println(v)
	v.X = 4 //we can also access the inner variables
	fmt.Println(v.X)
	fmt.Println(v.Y)

	//other notations:
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	p  := &Vertex{1, 2} // has type *Vertex (pointer)
	fmt.Println(p, v2, v3)
}

