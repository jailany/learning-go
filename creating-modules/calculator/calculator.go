package calculator

import "fmt"

func Add(x int, y int) int {
	fmt.Sprintf("The numbers are %v and %v", x, y)
	result := x + y
	return result
}