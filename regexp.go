package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`^f[a-zA-Z]*n$`)

	fmt.Println(regex.MatchString("firman"))
	fmt.Println(regex.MatchString("f1rmin"))
	fmt.Println(regex.MatchString("firman"))

	fmt.Println(regex.FindAllString("firman firmen firmoo farmon f1rman f1rm4n firman", -1))
}
