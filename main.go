package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Server is running")
	concatStr := fmt.Sprintf("Duong", "Quan")
	log.Printf("%s", concatStr)
	bol := true
	fmt.Println(!bol)
}
