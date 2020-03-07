package main
import (
    "fmt"
)


func main() {
    //String slice
    icecreams := []string{"chocolate", "coffee", "mango", "pista"}
    fmt.Printf("Ice Creams : %v Type : %T\n", icecreams, icecreams)
    //Index
    fmt.Printf("Value at index :  %v\n", icecreams[1])
    //Slicing a slice
    fmt.Printf("Partial Slice : %v\n",icecreams[1:])

    //Looping demo

    fmt.Print("Looping options\n")
    fmt.Println("Vanilla For loop")
    for i := 0; i < len(icecreams); i++ {
        fmt.Printf("Item at index %d is %s\n",i, icecreams[i])
    }
    fmt.Println("Using Range")
    for i := range icecreams {
        fmt.Printf("Item at index %d is %s\n",i, icecreams[i])
    }

    fmt.Println("Using Range for both index and content : Double value range")
    for i,name := range icecreams {
        fmt.Printf("Item at index %d is %s\n",i, name)
    }
    fmt.Println("Using Range for both index and content : ignoring index with _")
    for _,name := range icecreams {
        fmt.Printf("Item :  %s\n", name)
    }
}
