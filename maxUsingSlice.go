package main
import (
    "fmt"
)


func main() {
    //Integer Slice
    nos := [] int{2, 3, 23, 10, 34, 56, 60 }
    max := nos[0]

    for _, no := range nos[1:] {
        if max < no {
            max = no
        }
    }
    fmt.Println("Max no. :", max)
}
