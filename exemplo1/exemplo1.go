package exemplo1

import "errors"

var (
	ErrInputIsZero = errors.New("input can't be zero")
)

func isOdd(num int) (bool, error) {
	if num == 0 {
		return false, ErrInputIsZero
	}
	return num%2 == 0, nil
}
