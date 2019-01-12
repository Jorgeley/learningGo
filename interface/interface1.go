package main

import (
	"fmt"
	"math"
)

//the interface is just a type within methods signatures
type Abser interface {
	Abs() float64
}

type MyFloat float64

//this is basically the implementation of the interface
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {//this is very confusing for OOP programmers, but remember, Go is not OOP
	var abser Abser
	abser = MyFloat(-math.Sqrt2)
	fmt.Println(abser.Abs())
}

