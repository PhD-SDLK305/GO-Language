package main

func rename(name *string) {
	*name = "quan new"
}

func main() {
	a := 1
	b := &a
	a = 2
	c := &b
	println(*b)
	print(**c)
	name := "quan"
	rename(&name)
	println(name)
}
