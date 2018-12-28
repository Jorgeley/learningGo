package main

import "fmt"

//initializing variables
var i, j int = 1, 2

func main() {
	/* initializing with different types values
	 * notice that the type is ommited, since the variable type
	 * will be automatically set by the value type
	 */
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

