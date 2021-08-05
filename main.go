package main

import (
	"log"
)

func Test(secret string) {
	log.Println("THIS IS INSDIE A MISTAKE", secret)
}
func main() {
	log.Println("HELLO")
	var secret = "aftasvqweqweqwd"
	Test(secret)
}
