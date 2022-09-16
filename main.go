package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Golang programming tour", "GuideBook")
	flag.StringVar(&name, "n", "Golang programming tour", "GuideBook")
	flag.Parse()

	log.Printf("name: %s", name)
}
