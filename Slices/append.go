package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	var todo  []string

	todo = append(todo, "sing")
	todo = append(todo, "sing", "sang","song")


	s.Show("todo", todo)
}