package main

import (
	"fmt"
)

func main() {
	s := "gopher"
	fmt.Println("Hello and welcome, %s!", s)

	for i := 1; i <= 7; i++ {
		fmt.Println("i =", 100/i)
	}
}
