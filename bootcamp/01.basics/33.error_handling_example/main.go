package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Printf("SUCCESS: Converted %q to %d.\n", age, n)
}

// go run main.go 22 output
// SUCCESS: Converted "22" to 22.
// go run main.go 2a output
// ERROR strconv.Atoi: parsing "2a": invalid syntax
