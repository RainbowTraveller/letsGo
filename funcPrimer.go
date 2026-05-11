package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divMod(a int, b int) (int, int) {
	return a / b, a % b
}

func concat(a string, b string) string {
	return a + b
}

func noYearsUntil(age int) (int, int, int) {
	yearUntilAdult := 18 - age
	yearUnitilDrink := 21 - age
	yearsUnitilDrive := 16 - age

	return yearUntilAdult, yearUnitilDrink, yearsUnitilDrive
}

func main() {

	fmt.Printf("Simple Addition function : %d ", add(5, 3))
	fmt.Println()
	div, mod := divMod(5, 3)
	fmt.Printf("Return 2 values: Division : %d, Mod : %d", div, mod)
	fmt.Println()
	fmt.Printf("Concatenating 2 strings : %s ", concat("Hello ", "World"))

	var age int
	fmt.Println()
	fmt.Println("Enter your age : ")
	fmt.Scanf("%d", &age)
	adult, drink, drive := noYearsUntil(age)
	fmt.Printf("Years until you can drink : %d, drive : %d, be an adult : %d\n", drink, drive, adult)
}
