/**
 *Filename: pi_test.go
 * Usage: Tests the pi.go method.
 * Author: Nyah Check
 */
package pi

import (
	"fmt"
	"testing"

	"github.com/Ch3ck/NumGo/pi"
)

func TestPi(t *testing.T) {
	for i := 0; i < 10; i++ {
		value := pi.Pi(i)
		if value != 3 {
			t.Errorf("PI(%d) == %f which is wrong", i, value)
		} else {
			fmt.Println(pi.Pi(i))
		}
	}
}
