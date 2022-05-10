package main
import (
	"fmt"
    "os"
)

func main() {
    // i := 1
    // fmt.Printf("%q\n", os.Args[i])

    // i++
    // fmt.Printf("%q\n", os.Args[i])

    // i++
    // fmt.Printf("%q\n", os.Args[i])

    for i := 1; i < len(os.Args); i++ {
        fmt.Printf("%q\n", os.Args[i])
    }
}