package main

import (
	"fmt"
)

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

func main() {
	 prev := [3] string{
	  "Kafka Ravannge",
	  "Stay Golden",
      "Everthingship",
	}

	var books [4]string



	for i, b:= range prev{
		books[i] +=  b + "2nd Ed"
	}

	books[3] = "Neverland"

	fmt.Printf("Last year:\n%#v\n", prev)
	fmt.Printf("\nThis year:\n%#v\n", books)

	// var(
	// 	wBooks [winter] string
	// 	sBooks [summer] string
	// )


	// wBooks[0] = books[0]

	// for i := 0; i < len(sBooks); i++ {
	// 	sBooks[i] = books[i+1]
	// }

	// // for _, v:= range sBooks {
	// // 	fmt.Println(v)
	// // }

	// fmt.Printf("\nsummer: %#v\n", sBooks)
	// fmt.Printf("\nwinter: %#v\n", wBooks)

	// var publisher[len(books)] bool

	// publisher[0] = true
	// publisher[len(books)-1] = true
	// fmt.Printf("\npublisher book:")
	// for i, ok := range publisher {
	// 	if ok {
	// 		fmt.Printf("+ %s\n", books[i])
	// 	}
	// }
}