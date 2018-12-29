package main

import "fmt"

func adder() func(int) int {//closures are functions which returns functions
	sum := 0
	return func(x int) int { //this function is enclosed by adder() function
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder() //they are the function 'func(int) int'
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), //this will execute the 'func(x int) int' and return it
			neg(-2*i), //the same
		)
	}
}

