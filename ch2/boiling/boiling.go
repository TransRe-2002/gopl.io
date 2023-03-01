// Boiling prints the boiling point of water.
package main

import "fmt"

const boildingF = 212.0

func main() {
	var f = boildingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
