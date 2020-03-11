package main

import (
	"fmt"
)

func main() {
	x := 2

	//No brace around the condition
	if x > 10 {
		fmt.Println("X is bigger than 10")
	}

	if x > 100 {
		fmt.Println("X is very big")
	} else {
		fmt.Println("X is not big")
	}

	if x > 5 && x < 15 {
		fmt.Println("X is just right")
	}

	if x < 20 || x > 40 {
		fmt.Println("x is not in the range")
	}

	a, b := 1, 2

	//Value can be computed before the condition
	if add := a + b; add < 10 {
		fmt.Println("Addition and condition")
	}

	//Switch has 2 versions
	//No break is required

	//Classic switch, no break needed
	switch x {
	case 1:
		fmt.Println(" Case 1")
	case 2:
		fmt.Println(" Case 2")
	case 3:
		fmt.Println(" Case 3")
	default:
		fmt.Println(" Case Default")
	}

	//Naked switch where conditions are explicitly specified
	switch {
	case x > 100:
		fmt.Println("x > 100")
	case x < 100:
		fmt.Println("x < 100")
	}
}
