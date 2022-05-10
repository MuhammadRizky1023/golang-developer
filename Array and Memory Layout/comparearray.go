package main
import (
	"fmt"
)

func main()  {
	var(
		blue = [3]int{6, 9, 3}
		red = [3]int{6, 9, 3}
	)
	fmt.Printf("blue books: %v\n", blue)
	fmt.Printf("red books: %v\n", red)
	fmt.Printf("Are the equel?", blue == red)
}