package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	//functions can be assigned to variables
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) //functions can be passed as arguments

	fmt.Println(compute(hypot)) //even nested arguments
	fmt.Println(compute(math.Pow))
}

