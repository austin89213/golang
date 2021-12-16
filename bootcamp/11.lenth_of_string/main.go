package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "飛資得"
	fmt.Println(len(name))
	// output : 9
	fmt.Println(utf8.RuneCountInString(name))
	// output : 3
}
