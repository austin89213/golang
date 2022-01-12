package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st artgument:", os.Args[1])
	fmt.Println("2nd artgument:", os.Args[2])
	fmt.Println("3rd artgument:", os.Args[3])
	fmt.Println("Numbers of items inside of os.Args:", len(os.Args))
}
