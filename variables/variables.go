package main

import "log"

func main() {
	var (
		name    string
		thontin string
	)
	name = "quan"
	thontin = "hai duong"

	var str string = "str"
	var str2 = "str2"
	str3 := "str3"
	bol := true
	var a, b, c, d int
	a = 1
	b = 2
	c = 3
	d = 4

	e, f, g := 5, 6, 7
	log.Printf("%s%s%s%s", a, b, c, d)
	log.Printf("%s%s%s", e, f, g)
	log.Printf("%s%s", name, thontin)
	log.Printf("%s", str)
	log.Printf("%s", str2)
	log.Printf("%s", str3)
	log.Printf("%s", bol)
}
