/**
 *Filename: pi.go
 * Usage: Estimates PI to the nth digit using the  Bailey–Borwein–Plouffe formula (BBP formula).
 * Author: Nyah Check
 */
package NumGo

import (
	"fmt"

	"github.com/Ch3ck/NumGo/pi"
)

func main() {
	var n int

	fmt.Println("This program estimates the value of PI to the Nth decimal place")
	fmt.Printf("Please Enter n: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Value of PI to %d decimal Place: ", n)
	fmt.Println(pi.Pi(n))
}
