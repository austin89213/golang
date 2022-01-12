package main

import "fmt"

// func main() {
// 	var brand = "Google"
// fmt.Printf("%q\n", brand)
// output : "Google"
// 	fmt.Printf("%s\n", brand)
// output : Goolge
// }

// example
func main() {
	ops := 300
	ok := 200
	fail := 100
	fmt.Println("total:", ops, "success:", ok, "/", fail)
	fmt.Printf("total: %d success: %d / %d\n", ops, ok, fail)
}
