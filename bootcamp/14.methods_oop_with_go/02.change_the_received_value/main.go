package main

func main() {
	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "mirecraft",
		price: 20,
	}
	tetris := game{
		title: "tetris",
		price: 5,
	}
	mobydick.print()
	minecraft.discount(.5)

	minecraft.print()
	tetris.print()

	var h huge
	for i := 0; i <= 10; i++ {
		h.addr()
	}
}
