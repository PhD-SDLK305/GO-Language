package main

import "fmt"

func switchCase() {
	expression := 1
	switch expression {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("default")
	}
}
