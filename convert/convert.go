package main

import "strconv"

func main() {
	a, err := strconv.Atoi("1")
	if err != nil {
		println(err)
	} else {
		println(a)
	}

	stra := strconv.Itoa(3456)
	println(stra)
}
