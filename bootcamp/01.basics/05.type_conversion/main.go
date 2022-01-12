package main

import "fmt"

func main() {
	speed := 100 // speed is int
	force := 2.5 // force is float
	// speed = speed * force //this won't work
	// speed = speed * int(force) // this will get the wrong result which is 200 not 250
	speed = int(float64(speed) * force) //
	fmt.Println(speed)
	fmt.Println(force, int(force))
}
