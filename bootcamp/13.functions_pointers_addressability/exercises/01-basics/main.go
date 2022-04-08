// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var null *computer
	// compare the pointer variable to a nil value, and say it's nil
	if null == nil {
		fmt.Println("It's nil!")
	}
	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "apple"}
	// put the apple into a new pointer variable
	newApple := apple
	// compare the apples: if they are equal say so and print their addresses
	if apple == newApple {
		fmt.Printf("apple and newApple are equal\napple address: %p, newApple address: %p\n", apple, newApple)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "sony"}
	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if apple != sony {
		fmt.Printf("apple and sony are not equal\napple address: %p, sony address: %p\n", apple, sony)
	}
	// put apple's value into a new ordinary variable
	appleVal := *apple
	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple pointer addresses:%p, it points to:%p, appleVal address:%p\n", &apple, apple, &appleVal)
	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	// print the values:
	// the value that is pointed by the apple and the new variable
	if *apple == appleVal {
		fmt.Printf("They are equal: \napple point to value:%+v appleVal value:%+v\n", *apple, appleVal)
	}
	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	fmt.Println(sony.brand)
	// print the returned value
	fmt.Printf("appleVal: %+v\n", valueOf(apple))

	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))

}

// create a new function: change
// the func can change the given computer's brand to another brand
func change(c *computer, b string) {
	c.brand = b
}

// write a func that returns the value that is pointed by the given *computer
func valueOf(c *computer) computer {
	return *c
}

// write a new constructor func that returns a pointer to a computer
func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
