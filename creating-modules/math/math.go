package main

import (
	"fmt"
	"log"
	"module.com/calculator"
)

func main() {

	log.SetPrefix("Logger: ")
	log.SetFlags(0)

	result, err := calculator.Add(8, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}