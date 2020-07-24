package main

import (
	"fmt"
)

type Result struct {
	result int
}

func Test(result *Result) {
	result.result = 1
	fmt.Println("test")
}

func main() {
	result := new(Result)

	Test(result)

	fmt.Println(result.result)
}
