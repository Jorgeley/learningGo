package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//also, 'switch' constructor also can have the short statement
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("We're all gonna die!")
	default:
		fmt.Printf("%s.", os)
	}

	//we can also have a naked switch, useful for long if-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}

