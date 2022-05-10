package main
import (
	"fmt"
)

func main() {
	// var(
	// 	sum int
	// )

	// for i := 1; i < 50; i++ {
	// 	sum+= i
	// }
	// fmt.Println(sum)

	var(
		sum int
		i = 1
	)

	for  {
		if i > 10 {
			break
		}
		if i%2 != 0 {
			i++
			continue
		}
		sum += i
		fmt.Println( i, "-->", sum )
		i++
	}
	fmt.Println(sum)


	
}