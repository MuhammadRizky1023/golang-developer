package main
import (
	"os"
	"fmt"
	"time"
	"math/rand"
)
func main()  {
	args :=  os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[Your name] [Positive][Negative]")
		return
	}
	name, mood := args[0], args[1]
	moods := [...][3]string{
		{"happy😁","good👍", "awesome😎"}, 
	    {"sad😞", "bad👎", "terrible🙁"},
	}


	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods[0]))

	var mi int
	if mood !=  "Positive"{
		mi = 1
	}


	fmt.Printf("%s feel %s\n",  name, moods[mi][n])

}