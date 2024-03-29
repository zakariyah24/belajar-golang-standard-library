package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	}

	convStr := "1000"
	resultInt, err := strconv.Atoi(convStr) // str to int
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	resultStr := strconv.Itoa(10009) // int to str
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultStr)
	}
}
