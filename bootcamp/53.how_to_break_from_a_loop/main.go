package main

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}
		sum += i
		i++
	}
}
