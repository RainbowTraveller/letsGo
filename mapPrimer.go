package main

import (
	"fmt"
)

func main() {
	capitals := map[string]string{
		"India": "Delhi",
		"USA":   "DC",
		"Japan": "Tokyo",
	}

	fmt.Print("No. of items in the Map")
	fmt.Println(len(capitals))

	//Getting value for a key
	fmt.Printf("Key : %s Value : %s\n", "India", capitals["India"])

	//Get 0 if key is not present
	fmt.Printf("Key : %s, Value : %s\n", "France", capitals["France"])

	//Get pair of responses
	value, ok := capitals["USA"]
	if !ok {
		fmt.Println("Value not found for given key")
	} else {
		fmt.Println(value)
	}

	//Adding a pair to existing map
	capitals["France"] = "Paris"

	fmt.Println(capitals)
	fmt.Println("Iterating......")
	fmt.Println("Only Keys Iterating......")

	for key := range capitals {
		fmt.Printf("(%s,%s), ", key, capitals[key])
	}

	fmt.Println()
	fmt.Println("Keys and Values Iterating......")
	for key, value := range capitals {
		fmt.Printf("(%s,%s), ", key, value)
	}
	fmt.Println()
	fmt.Println("Deleting")
	delete(capitals, "India")
	fmt.Println(capitals)
}
