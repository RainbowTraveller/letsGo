package main

import (
	"fmt"
	"os"
)

func main() {
	//All arguments
	argsWithProg := os.Args
	for _, ag := range argsWithProg {
		//First argument is the actual path of the  current program
		fmt.Println(ag)
	}
}
