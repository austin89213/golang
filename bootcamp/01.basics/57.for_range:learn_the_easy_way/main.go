package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := strings.Fields("lazy cat jumps again and again and again")
	// fmt.Println(len(words))
	for _, v := range words {
		fmt.Println(v)
	}
	for _, v := range os.Args[1:] {
		fmt.Printf("%q\n", v)
	}
	// 如果只想要 index, 可以把 v 拿掉而無須用 _ 代替的
	for i := range os.Args[1:] {
		fmt.Printf("%d\n", i)
	}
}
