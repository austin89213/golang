package main

import "fmt"

func main() {
	score, valid := 4, false
	if score > 3 && valid {
		fmt.Println("GOOD!")
	} else if score == 3 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("NOOOO")
	}
}
