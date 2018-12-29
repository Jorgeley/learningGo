package main

import (
	"fmt"
	"strings"
)

func main() {
	var a[2]string //an array of 2 strings
	a[0] = "Hello" //assignment at 1st position
	a[1] = "World" //assignment at 2nd position
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//we can also assign an array within values
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	/*
	 * IMPORTANT: the array size is fixed, we can't do this:
	 * primes[7] = 99
	 * it results: invalid array index 7 (out of bounds for 6-element array)
	 */

	 //in the other hand, we can do this to have a dinamically sized array:
	 b := []int{99, 22, 11, 66}
	 fmt.Println(b)
	 //we can also get slices of the array using the syntax array[init:end]
	 var c[]int = b[1:3]
	 fmt.Println(c)
	 //be careful with slices, they refer to the pointer, so if slice changed, the original array changes too:
	 c[0] = 44 //b is affected too
	 fmt.Println(b)
	 //we can also ommit the init and end of the slice:
	 d := b[2:] //from 2 to the rest
	 fmt.Println(d)
	 e := b[:3] //from 0 to 3
	 fmt.Println(e)
	 f := b[:] //all elements
	 fmt.Println(f)
	 //we can use 'make' to create slices too
	 g := make([]int, 0, 5)
	 fmt.Printf("len=%d cap=%d %v\n", len(g), cap(g), g)
	 //we can add more elements:
	 h := append(g, 99, 66, 1, 44, 77, 33, 2, 4, 7, 0, 5, 221, 123, 77, 896)
	 fmt.Printf("len=%d cap=%d %v\n", len(h), cap(h), h)

	 //this are slices of slices:
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	//this is more used and easier to iterate an array:
	for i, v := range board{
		for ii, vv := range v{
			fmt.Println(i, ii, vv)
		}
		//we can skip the index:
		for _, vvv := range v{
			fmt.Println(vvv)
		}
		//or we can skip the value:
		for iii := range v{
			fmt.Println(iii)
		}
	}

	 //this is cool, we can create an array of structures:
	 s := []struct { //here we define that 's' is an array structure
		i int
		b bool
	}{ //and here we define the values of the structure
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

