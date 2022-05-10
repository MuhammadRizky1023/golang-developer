package main
import (
	"strconv"
	"fmt"
	"os"
)

func main() {
	age := os.Args[1]

	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("Succes convert  %q to %d\n", age, n)
}