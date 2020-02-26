package main

import (
	"fmt"
)

func main() {

	//No braces and no pre increment
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------")
	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("-------")
	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("-------")
	//with only condition
	i := 0
	for i < 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-------")
	//Just like while(true)
	for {
		fmt.Println(i)
		i++
		if i > 3 {
			break
		}
	}
}
