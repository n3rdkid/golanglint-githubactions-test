package main

import (
	"log"
)

func Test() {
	log.Println("THIS IS INSDIE A MISTAKE")
}
func main() {
	log.Println("HELLO")
	var password = "aftasvqweqweqwd" // nolint:staticcheck
	var abc = 2                      // nolint:staticcheck
	Test()
}
