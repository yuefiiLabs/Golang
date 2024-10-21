package main

import "fmt"

func main() {
	// Slice
	var slice []int = []int{1, 2, 3}

	// Menambahkan elemen
	slice = append(slice, 4)

	fmt.Println("Slice:", slice)
}
