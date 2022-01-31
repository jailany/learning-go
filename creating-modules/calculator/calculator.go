package calculator

import (
	"fmt"
	"errors"
)

func Add(x int, y int) (int, error) {
	if (x == 0 || y == 0) {
		return 0, errors.New("The values passed cannot be 0")
	}
	fmt.Sprintf("The numbers are %v and %v", x, y)
	result := x + y
	return result, nil
}