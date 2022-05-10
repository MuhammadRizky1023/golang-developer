package main
import (
	"fmt"
	"os"
)
const (
	user = "Usage: [username] [password]"
	errorUser = "Access denied for %q.\n"
	errorPassword = "invalid passsword for %q.\n"
	completed = "Access ganted to %q.\n"

	username1, username2 = "roger", "luffy"
	password1, password2 = "3990", "1440"

)
func main() {
	args := os.Args
	if len(args) !=3 {
		fmt.Println(user)
		return
	}

	u, p := args[1], args[2]
	if u != username1 && u != username2 {
		fmt.Printf(errorUser, u)
	} else if u == username1 && p == password1  {
		fmt.Printf(completed, u)
	} else if   u == username2 && p == password2 {
		fmt.Printf(completed, u)
	} else {
		fmt.Printf(errorPassword, u)
	}

}