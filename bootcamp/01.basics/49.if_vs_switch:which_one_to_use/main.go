package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}
	// m := os.Args[1]
	// if m == "Dec" || m == "Jan" || m == "Feb" {
	// 	fmt.Println("Winter")
	// } else if m == "Mar" || m == "Apr" || m == "May" {
	// 	fmt.Println("Spring")
	// } else if m == "Jun" || m == "Jul" || m == "Aug" {
	// 	fmt.Println("Spring")
	// } else if m == "Sep" || m == "Oct" || m == "Nov" {
	// 	fmt.Println("Fall")
	// } else {
	// 	fmt.Printf("%q is not a month\n", os.Args[1])
	// }
	switch m := os.Args[1]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Spring")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a month\n", os.Args[1])

	}
}
