/**
 *Filename: Fibonacci_test.go
 * Usage: Estimates Nth Fibonacci series for any number.
 * Author: Nyah Check
 */
package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {

	for i := 0; i < 10; i++ {
		value, _ := fibonacci(i)
		if len(value) < 2 {
			t.Errorf("Fibonacci(%d) does not exist.", i)
		} else {
			fmt.Println(fibonacci(i))
		}
	}
}
