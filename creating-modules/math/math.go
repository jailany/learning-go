package main

import (
	"fmt"
	"log"
	"module.com/calculator"
)

func main() {

	log.SetPrefix("Logger: ")
	log.SetFlags(0)

	result, err := calculator.Add(8, 2)

	even, evenErr := calculator.GetEvenNumbers(10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	if evenErr != nil {
		log.Fatal(evenErr)
	}
	fmt.Println("The even numbers are:")
	fmt.Println(even)
}