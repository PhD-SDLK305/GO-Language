package main

import "fmt"

func main() {
	a := 1
	b := 2
	if !(a > b) {
		i := 1
		for i < 10 {
			if i == 5 {
				break
			}
			i++
		}
		fmt.Println("a > b")
	} else {
		fmt.Print("a < b")
	}
	switchCase()
}
