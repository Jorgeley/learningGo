package main

import (
	"fmt"
	"math"
)

/*Go doesn't have classes, but we can define types within methods (functions) like: type.function()
this is the 'method receiver'*/
type Vertex struct { //this is a new type
	X, Y float64
}

/*this is the method within the type Vertex, don't confuse with functions, functions are declared in
the sample pattern: func(x string) string
      receiver     method    return type*/
func (v Vertex)    Abs()     float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4} //declaring the new type
	fmt.Println(v.Abs()) //calling the type method
}

