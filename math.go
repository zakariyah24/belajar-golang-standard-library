package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.40))
	fmt.Println(math.Round(1.60)) // pembulatan ke yg paling dekat
	fmt.Println(math.Max(10, 11))
	fmt.Println(math.Min(10, 11))
}
