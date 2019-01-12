package main

import "fmt"

/*functions can return multiple results.
notice the return types within the 2nd brackets
           arguments        return*/
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	/*
	 * Outside a function, every statement begins with a 
	 * keyword (var, func, and so on) and so the := construct is not available.
	 * the short assignment ':=' constructor replaces 'var' declaration with types
	 */
	a, b := swap("hello", "world") //notice that 'var' declaration and types are missing since they are implicit
	fmt.Println(a, b)
}

