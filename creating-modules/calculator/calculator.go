package calculator

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Add(x int, y int) (int, error) {
	if (x == 0 || y == 0) {
		return 0, errors.New("The values passed cannot be 0")
	}
	fmt.Sprintf("The numbers are %v and %v", x, y)
	result := x + y + getARandomNumber()
	return result, nil
}

func getARandomNumber() (int) {

	evenNumbers := []int{
		2, 4, 6, 8, 10,
	}

	return evenNumbers[rand.Intn(len(evenNumbers))]
}