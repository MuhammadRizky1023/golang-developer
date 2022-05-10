package main
import  (
	"fmt"
	"os"
   "strings"
)

func main() {
	// fmt.Println("Code")

	msg :=  os.Args[1]
	l := len(msg)

	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)
	fmt.Println(s)
}