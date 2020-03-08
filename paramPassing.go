package main
import (
    "fmt"
)

func doublesAt(values [] int, index int) {
    values[index] *= 2
}

func double(val *int) {
    *val *= 2
}
func main() {
    values := [] int {1,2,3,4}
    doublesAt(values, 2)
    fmt.Println("Maps and Slices are passed by values")
    fmt.Println(values)

    fmt.Println()
    fmt.Println("Primitives are passed by value so use pointers")
    val := 2
    double(&val)
    fmt.Printf("Doubled value of premitive: %d\n", val)

}
