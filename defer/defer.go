package main

import "fmt"

/* 
 * defer is used as the same as 'finally' in try-catch blocks in other languages
 * the main purpose is to guarantee that a code is executed no matter what happens,
 * per example, when we need to close resources
 */

func main() {
	defer fmt.Println("world") //this will be deferred/delayed until next return
	fmt.Println("hello")

	//it's important to notice that the defered calls are executed from the last to the first:
	for i := 0; i < 10; i++ {
		defer fmt.Println("deferred ", i)
	}
}

