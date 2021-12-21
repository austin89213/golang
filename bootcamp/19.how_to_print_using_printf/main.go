package main

import "fmt"

// func main() {
// 	var speed int
// 	var heat float64
// 	var off bool
// 	var brand string

// 	fmt.Printf("%T\n", speed)
// 	fmt.Printf("%T\n", heat)
// 	fmt.Printf("%T\n", off)
// 	fmt.Printf("%T\n", brand)
// }

// %v - any value
// %q - string with quotes
// %s - string
// %d - int
// %f - float
// %t - bool
// %T - type of variable
// use index to decide which variable to refer ex: %[1]v, %[3]v
func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 244.701
		hasLife  = false
	)
	fmt.Printf("Planet: %v\n", planet)
	// output Planet: venus
	fmt.Printf("Distance: %v millions kms\n", distance)
	// output Distance: 261 millions kms
	fmt.Printf("Orbital Period: %v days\n", orbital)
	// output Orbital Period: 244.701 days
	fmt.Printf("Dose %v has life?: %v\n", planet, hasLife)
	// output Dose venus has life?: false
	fmt.Printf("%[2]v is %[1]v millions kms away. Think! \n", distance, planet)
	// output venus is 261 millions kms away. Think!
}
