package main

func main() {

	game := Game{
		View: NewVew(),
	}

	game.Start(&HelloScreen{})
}
