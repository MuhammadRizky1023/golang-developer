package main
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2{
		fmt.Println("Give your month")
		return
	}
	// month := os.Args[1]

	// if month == "Desember" || month == "January" || month ==  "February"{
	// 	fmt.Println("winter")
	// } else if  month == "March" || month == "April" || month ==  "Mei"{
	// 	fmt.Println("Spring")
	// }else if  month == "September" || month == "October" || month ==  "November"{
	// 	fmt.Println("Fall")
	// } else {
	// 	fmt.Printf("%q is not a month.\n", month)
	// }

	switch month := os.Args[1]; month{
	case "Desember" , "January" , "February":
			fmt.Println("winter")
	case  "March" , "April" , "Mei":
			fmt.Println("Spring")
	case  "September" , "October" ,"November":
			fmt.Println("Fall")
   default:
			fmt.Printf("%q is not a month.\n", month)
   }
}