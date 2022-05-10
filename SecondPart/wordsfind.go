package main
import (
	"os"
	"strings"
	"fmt"
)

const corpus = "" + "lazy dog jump again and again and again."

func main() {
	word := strings.Fields(corpus)
	query := os.Args[1:]


queries:
	for _, q := range query {
	search: 
		for i, w := range word {

			switch q{
			case "and", "or","the":
				break search
			}

			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				continue queries
			}
		}
	}
}