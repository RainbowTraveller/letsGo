package main

import (
	"fmt"
)

func main() {
	str := "This is the string"

	fmt.Println(len(str))

	fmt.Println(str[0])

    fmt.Printf("book[0] : %v (type %T)\n", str[0], str[0])

    //str[0] = 'c' // String are immutable

    //slicing
    //(start , end)
    fmt.Println("Slicing start to end : ", str[4:11])
    //no end
    fmt.Println("Slicing no end : ", str[4:])
    //no start
    fmt.Println("Slicing no start : ", str[:4])
    //Concatenate using +
    fmt.Println("This String " + " That string")
    //Unicode
    fmt.Println("Some Unicode string")
    //Multi Line
    multiline := `
        this is multiline
        String
    `
    fmt.Println(multiline)

}
