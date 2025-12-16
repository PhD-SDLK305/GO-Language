package main

import "fmt"

func main() {
	for i := 0; i < 5; i = i + 1 {
		fmt.Println(i)
	}
	j := 0
	for j < 10 {
		if j == 5 {
			break
		}
		// if j != 5 {
		// 	continue
		// }
		fmt.Println(j)
		j++
	}
}
