package main

import "fmt"

func main() {

	switch i := 10; {
	//above line equals to switch i := 10; true
	//Since the true is just hidden, we still need to keep the ";" the seperate the statement and condition
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}
