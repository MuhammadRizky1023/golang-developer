package main
import "fmt"

func main() {
	var i int
loop:
     if i < 5 {
		 fmt.Println("looping")
		 i++
		 goto loop
	 }
	 fmt.Println("done")
}