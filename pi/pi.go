/**
 *Filename: pi.go
 * Usage: Estimates PI to the nth digit using the  Bailey–Borwein–Plouffe formula (BBP formula).
 * Author: Nyah Check
 */
package pi

import "errors"

// pi returns the value of pi to the nth decimal.
//
// n >= 0
//Special cases are:
//	pi(-n) = p(n)
//	pi(NaN) = pi(0)
func Pi(n int) (float64, error) {
	var result float64 = 0
	if n >= 0 {
		for i := 0; i < n; i++ {
			result += (1.0 / 16.0 * float64(i)) * ((4.0 / (8.0*float64(i) + 1)) + (2.0 / (8.0*float64(i) + 4.0)) - (1.0 / (8.0*float64(i) + 5.0)) - (1.0 / (8.0*float64(i) + 6.0)))
		}
	} else {
		return result, errors.New("Invalid decimal number")
	}

	return result, nil
}
