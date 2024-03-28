package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Firman"))
	fmt.Println(slices.Index(names, "Firman"))
	fmt.Println(slices.Index(names, "George"))
	// its must have incresing order so the bottom of this return error
	n, found := slices.BinarySearch(names, "Paul")
	fmt.Println("Paul:", n, found)
	// fmt.Println(slices.BinarySearch(names, "Paul"))
}
