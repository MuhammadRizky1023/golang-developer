package main
import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)

const MaxTurn = 5

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	guest, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("This is not a number")
		return
	}

	for turn := 0; turn < MaxTurn; turn++ {
		n := rand.Intn(guest + 1)

		if n == guest {
			fmt.Println("🏆 YOU WIN")
			return
		}
	}
	fmt.Println("☠️ YOU LOST TRY AGAIN ...")
}