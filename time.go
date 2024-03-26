package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	parse, _ := time.Parse(formatter, "2024-03-04 15:46:00")
	fmt.Println(parse)
	fmt.Println(parse.Local())
}
