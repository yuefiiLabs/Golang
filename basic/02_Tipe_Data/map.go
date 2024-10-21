package main

import "fmt"

func main() {
	// Map
	var m map[string]int = make(map[string]int)
	m["Alice"] = 25
	m["Bob"] = 30

	fmt.Println("Map:", m)
}
