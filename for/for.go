package main

import "fmt"

/*Go has only one loop constructor which is the 'for' bellow*/

func main() {
	sum := 0
	//this is the complete structure:
	//   init      condition     iteration
	for i := 0;     i < 10;        i++ {
		fmt.Println("iteration ", i)
		sum += i
	}
	fmt.Println(sum)

	s := 1
	//the init and iteration are optional
	for ; s < 1000; {
		fmt.Print(".")
		s += s
	}
	fmt.Println(s)

	t := 1
	//you can also drop the semicolons, in other words, this is a 'while' loop
	for t < 1000{
		fmt.Print(".")
		t += t
	}
	fmt.Println(t)

	//moreover, you can have a infinity loop, careful!
	for {
		fmt.Println("this is gonna be forever... (CTRL+C to quit")
	}

}
