package main

import (
	"color/functions"
	"fmt"
)

func main() {
	result := functions.Result()
	if result != nil {
		fmt.Println(result)
	}
}
