//Calculate the mean of 2 numbers
package main

import (
	"fmt"
)

func main() {
	/*var x int / float64
	var y int / float64
	x = 1.0
	y = 2.0
	*/

	x, y := 1.0, 2.0

	fmt.Printf("x= %v, Type = %T\n", x, x)
	fmt.Printf("y= %v, Type = %T\n", y, y)

	// var mean int / float64
	mean := (x + y) / 2.0
	fmt.Printf("mean= %v, Type = %T\n", mean, mean)
}
