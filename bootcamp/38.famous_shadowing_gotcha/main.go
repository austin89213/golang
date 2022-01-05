package main

import (
	"fmt"
	"os"
	"strconv"
)

// below will shadow u
// func main() {
//	var n int
//	if a := os.Args; len(a) != 2 {
//		only: a variable
//		fmt.Println("Give me a number.")
//	} else if n, err := strconv.Atoi(a[1]); err != nil {
//		only: a, n and err variable
//		fmt.Printf("Cannot convert %q .\n", a[1])
//	} else {
//		all the declared variables can be used only in the if statement
//		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
//	}
//	fmt.Printf("n is %d, you've been shadowed", n)
//}

// beloew is the corrcet way
func main() {
	var (
		n   int
		err error
	)
	if a := os.Args; len(a) != 2 {
		//only: a variable
		fmt.Println("Give me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		//only: a, n and err variable
		fmt.Printf("Cannot convert %q .\n", a[1])
	} else {
		//all the declared variables can be used only in the if statement
		n *= 2
		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}
	fmt.Printf("n is %d, you've been shadowed", n)
}
