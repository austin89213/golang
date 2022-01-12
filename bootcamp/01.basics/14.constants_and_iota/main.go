package main

import "fmt"

// func main() {
// 	// const (
// 	// 	monday    = 0
// 	// 	tuesday   = 1
// 	// 	wednesday = 2
// 	// 	thursday  = 3
// 	// 	friday    = 4
// 	// 	saturday  = 5
// 	// 	sunday    = 6
// 	// )
// 	const (
// 		monday = iota + 1
// 		tuesday
// 		wednesday
// 		thursday
// 		friday
// 		saturday
// 		sunday
// 	)
// 	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
// }

func main() {
	const (
		EST = -5
		MST = -7
		PST = -8
	)
}

// 可以改寫成以下

func main() {
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
	fmt.Println(EST, MST, PST)
}
