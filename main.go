package main

import "log"

func controller() {
	log.Println("test")
}

func main() {
	go controller()
	log.Println("test")
}
