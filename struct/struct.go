package main

type Employee struct {
	name    string
	address string
}

func showInfo(std *Employee) {
	println(*&std.name)
	println((*std).address)
	println(std.address) // dung struct co the bo * bien thi bat buoc *
}

// receiver
func (std *Employee) showInfo2() {
	println(*&std.name)
	println((std).address)
	println(std.address) // dung struct co the bo * bien thi bat buoc *
}

func (std *Employee) clear() {
	std.name = ""
	std.address = ""
}

func main() {
	quan := Employee{
		name:    "Duong Minh Quan",
		address: "hai duong",
	}
	quan.showInfo2()
	quan.clear()
	quan.showInfo2()
	showInfo(&quan)
}
