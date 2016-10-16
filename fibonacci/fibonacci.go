/**
 *Filename: Fibonacci.go
 * Usage: Estimates Nth Fibonacci series for any number.
 * Author: Nyah Check
 */
package fibonacci

import "errors"

// Fibonacci: Returns the Nth fibonacci sequence for any N
// N must be greater than 0
// Fibonacci(N)
func fibonacci(N int) ([]int, error) {
	sequence := make([]int, 0, 5)
	if N < 1 {
		return sequence, errors.New("Invalid N value for Sequence")
	} else {
		j := N - 1
		k := N
		for i := 0; i < N; i++ {
			sequence = append(sequence, j+k)
			j++
			k++
		}
	}

	return sequence, nil
}
