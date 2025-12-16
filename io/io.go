package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// var hoten string
	// fmt.Println("vui long nhap ho ten : ")
	// fmt.Scan(&hoten)
	// log.Printf("%s", hoten)
	var fullName string
	fmt.Println("vui long nhap ho ten : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fullName = scanner.Text()
	}
	log.Printf("%s", fullName)
}
