package main

import (
	"fmt"
	"math"
)

func SimpleEquations(a, b, c int) {
	// x + y + z = A
	// xyz = B
	// x^2 + y^2 + z^2 = C
	var x, y, z int
	var solution bool

	for x = 1; float64(x) <= math.Sqrt(float64(c)); x++ {
		for y = x; float64(y) <= math.Sqrt(float64(c)); y++ {
			z = a - x - y                       // x + y + z = A
			if x*y*z == b && x*x+y*y+z*z == c { // xyz = B dan x^2 + y^2 + z^2 = C
				solution = true
				break
			}
		}
		if solution {
			fmt.Println(x, y, z)
			break
		}
	}
	if !solution {
		fmt.Println("no solution")
	}
}
func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
