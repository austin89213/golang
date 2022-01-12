package main

import (
	"fmt"
	"strconv"
)

// func main() {
// 	n, err := strconv.Atoi("42")
// 	if err != nil {
// 		fmt.Println("There was an error, n is ", n)
// 	}
// }

func main() {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("There was an error, n is ", n)
	}
}
