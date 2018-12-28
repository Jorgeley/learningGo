package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	 * when importing packages, only the capital names are exported like 'Pi, Now'
	 * this will result an error 'cannot refer to unexported name math.pi'
	fmt.Println(math.pi)
	* because the 'pi' is not capital letter, so it's not exported, but this will work:
	*/
	fmt.Println(math.Pi)
}

