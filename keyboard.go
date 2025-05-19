// Package keyboard provides a function to convert an integer to a float64.
package keyboard

import (
	"strconv"
)

// Getdouble converts an integer to a float64
// and returns it. It takes an integer as input.
func GetFloat(input string) (float64, error) {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
