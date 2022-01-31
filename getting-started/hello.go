package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("Created the main package.")
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}