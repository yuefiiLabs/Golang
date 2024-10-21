package main

import "fmt"

// Struct
type Person struct {
	Name string
	Age  int
}

func main() {
	// Membuat objek Person
	p := Person{Name: "Alice", Age: 30}

	fmt.Println("Struct:", p)
}
