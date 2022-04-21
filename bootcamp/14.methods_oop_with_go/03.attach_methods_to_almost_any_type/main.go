package main

func main() {
	var (
		minecraft = game{
			title: "mirecraft",
			price: 20,
		}
		tetris = game{
			title: "tetris",
			price: 5,
		}
	)

	var items []*game
	items = append(items, &minecraft, &tetris)
	my := list(items)
	my.print()
}
