package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 244.701
		hasLife  = false
	)
	fmt.Printf("Planet: %d\n", planet)
	// output Planet: venus
	fmt.Printf("Distance: %d millions kms\n", distance)
	// output Distance: 261 millions kms
	fmt.Printf("Dose %s has life?: %t\n", planet, hasLife)
	// output Dose venus has life?: false
	fmt.Printf("%[2]v is %[1]v millions kms away. Think! \n", distance, planet)
	// output venus is 261 millions kms away. Think!

	// how %f works:
	fmt.Printf("Orbital Period: %f days\n", orbital)
	// output Orbital Period: 244.701000 days
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	// output Orbital Period: 245 days
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	// output Orbital Period: 244.7 days
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	// output Orbital Period: 244.70 days
	fmt.Printf("Orbital Period: %.3f days\n", orbital)
	// output Orbital Period: 244.701 days
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
	// output Orbital Period: 244.7010 days

}
