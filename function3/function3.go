package main

import "fmt"

/* Naked Return: an empty/naked 'return' always returns the 
 * named variables at the top of the function*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //it will return 'x' and 'y' declared at the top
}

func main() {
	fmt.Println(split(17))
}

