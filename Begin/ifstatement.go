package main
import (
	"fmt"
)
func main() {
	score := 95
	if score == 10 {
		fmt.Println(" Yu are so smart")
	} else if score >= 9 {
		fmt.Println("Your smart")
	} else if score >= 8 {
		fmt.Println("You are good are so good")
	} else if score  >= 7 {
		fmt.Printf("You are good")
	} else {
		fmt.Println("you are dumb")
	}
}