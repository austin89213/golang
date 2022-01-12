package main

import "fmt"

func main() {
	i := 142
	// switch {
	// case i > 100:
	// 	fmt.Println("big positive number")
	// case i > 0:
	// 	fmt.Println("positive number")
	// default:
	// 	fmt.Println("number")
	// }
	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number ")
	}
	fmt.Println()
}
