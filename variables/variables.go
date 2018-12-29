package main

import "fmt"

//'var' declares variables, if type is the same, only need to specify the last one
var c, python, java bool

/*Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func main() {
	/*
	Variables declared without an explicit initial value are given their zero value.
	The zero value is:
		0 for numeric types,
		false for the boolean type, and
		"" (the empty string) for strings.
	*/
	var i int
	fmt.Println(i, c, python, java)

	l = 3.12 //if no type specified, it's automatically inferred, in this case, float
}

