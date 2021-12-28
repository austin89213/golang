package main

import (
	"fmt"
	"os"
	"strconv"
)

// note: no need to escape ''
// use strconv.ParseFloat
// %g
func main() {
	v := os.Args[1]
	// feet, err := strconv.Atoi(v) 如果非整數會有問題
	feet, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Printf("error: '%v' is not a number.\n", v)
		return
	}
	meter := float64(feet) * 0.3048
	fmt.Printf("%g feet is %g meters\n", feet, meter)
}
