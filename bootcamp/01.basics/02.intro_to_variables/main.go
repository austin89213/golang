// Easy Sample
var speed int

// Mutiple Declarations
// You can declare multiple variables at once
// EX1:
func main() {
	var speed int
	var heat float64
	var off bool
	var brand string
}

// EX2:
func main() {
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)
}

// You can declare multiple variables using the type once
func main() {
	var speed, velocity int
}

// Let's say that you want to declare a variable and you already know its value
// Go can figure out the type automagically
func main() {
	var safe = true
	var speed = 100
}