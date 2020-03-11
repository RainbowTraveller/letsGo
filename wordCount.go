package main

import (
	"fmt"
	"strings"
)

func main() {
	tracker := map[string]int{}
	text :=
		`
    Needles and pins
    Needles and pins
    Sew me a sail
    To catch me the wind
    `
	fmt.Println(text)
	fmt.Printf("All fields : %q \n", strings.Fields(text))

	words := strings.Fields(text)
	for _, word := range words {
		tracker[strings.ToLower(word)]++
	}

	fmt.Println(tracker)
}
