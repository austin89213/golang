package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println(
		"dir:", dir,
		"file:", file)
}

//如果不想要用變數, 可用 _ 代替, 如下
// func main() {
// 	var file string
// 	_, file = path.Split("css/main.css")
// 	fmt.Println(
// 		"file:", file)
// }
