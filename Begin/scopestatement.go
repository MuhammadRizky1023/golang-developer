package main
import (
	"strconv"
	"fmt"
	"os"
)

func main() {
	var (
		n int
		err error
	)
	if a:= os.Args; len(a) != 2 {
		fmt.Println("Give yur number")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		n *= 2
		fmt.Printf("%s * 2  %d.\n", a[1], n*2)
	}

	fmt.Printf("n is %d. 👻👻👻👻 - you have been gosht", n)
}