package main

import "fmt"

//Attention! Type after the variable
func add(x int, y int) int {
	return x + y
}

//if the types are the same, we only need to specify the last one
func addd(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13), "\n")
	fmt.Println(addd(4, 3), "\n")
}

