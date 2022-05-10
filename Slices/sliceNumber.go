package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"strconv"
)


func main() {
	rand.Seed(time.Now().UnixNano())


	 max, _ := strconv.Atoi(os.Args[1])
	var uniques []int

	for len(uniques) < max {
		n:= rand.Intn(max) + 1
		fmt.Print(n," ")

loop:
		for _, u := range uniques{
			if  u == n {
				continue loop
			}
		}
		uniques = append(uniques, n)
	}

	fmt.Println("\n\nuniques", uniques)
}