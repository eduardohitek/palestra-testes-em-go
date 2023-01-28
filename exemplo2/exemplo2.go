package exemplo2

import "errors"

var (
	ErrInputIsZero         = errors.New("input can't be zero")
	ErrInputCantBeNegative = errors.New("input can't be negative")
)

func isOdd(num int) (bool, error) {
	if num == 0 {
		return false, ErrInputIsZero
	}
	if num < 0 {
		return false, ErrInputCantBeNegative
	}
	return num%2 == 0, nil
}
