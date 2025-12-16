package main

import "log"

const CCCD int32 = 123456789

func main() {
	const CCCD int32 = 123456789
	// const CCCD int32 = 123456789 X
	log.Printf("%s", CCCD)
}
