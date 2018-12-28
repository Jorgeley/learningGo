package main

import "fmt"

/*functions can return multiple results.
notice the return types within the brackets*/
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	//notice the attribution operator :=
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

