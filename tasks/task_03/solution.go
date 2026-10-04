package main

import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	if n == 0 {
		return "", ErrZero
	}

	var result string

	if n%3 == 0 {
		result += "Fizz"
	}

	if n%5 == 0 {
		result += "Buzz"
	}

	if len(result) != 0 {
		return result, nil
	}

	return strconv.Itoa(n), nil
}
