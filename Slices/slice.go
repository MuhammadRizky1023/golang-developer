package main
import (
	"fmt"
)

func main() {
	var books [5]string
	books[0] = "The dracula"
	books[1] = "1984"
	books[2] = "island"


	newBook := [5]string{"Fyodor", "Crime and Punishment"}
	books = newBook


	games := []string{"Zuma", "Pekemon"}
	newGames := []string{"Tetris", "Mortal Kombat", "PingPong"}


	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not"
			break
		}
	}

	if games == nil {
		ok = "not"
	}

	fmt.Printf("games and newgames: %sequel\n\n", ok)


	fmt.Printf("books: %v\n", books)
	fmt.Printf("games: %v\n", games)
	fmt.Printf("games: %T\n", games)
	fmt.Printf("games: %d\n", len(games))
}