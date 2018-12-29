package main

import "fmt"

func main(){

	x := true
	//this is the basic well known if structure
	if x{
		fmt.Println("'if' executed")
	}

	//and this is an 'if' within a short statement:
	//  statement    condition
	if y := true;       y != true{
		fmt.Println("another successful 'if' executed")
	}else{ //offcourse we can have 'else' as well
		fmt.Println("another not successful 'if' executed")
	}
}
