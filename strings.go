package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Zakariyah Firmansyah", "Zakariyah"))
	fmt.Println(strings.Split("Zakariyah Firmansyah", " "))
	fmt.Println(strings.ToLower("Zakariyah Firmansyah"))
	fmt.Println(strings.ToUpper("Zakariyah Firmansyah"))
	fmt.Println(strings.Trim("  Zakariyah Firmansyah      ", " "))
	fmt.Println(strings.ReplaceAll("Zakariyah Firmansyah", "Zakariyah", "Jaka"))
}
