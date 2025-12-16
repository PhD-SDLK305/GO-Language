package main

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("hello", name)
}

func returnTwoData(name string) (string, int) {
	return fmt.Sprintf("hello", name), 30
}

func main() {
	println(Hello("quan"))
	println(returnTwoData("quan"))
}
