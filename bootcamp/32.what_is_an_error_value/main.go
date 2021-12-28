// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	s := strconv.Itoa(42)
// 	fmt.Println(s)
// }

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number :", n)
	fmt.Println("Returned error value :", err)
}

// go run main.go 42 out put
// Converted number : 42
// Returned error value <nil>

// go run main.go A out put
// Converted number : 0
// Returned error value : strconv.Atoi: parsing "A": invalid syntax
