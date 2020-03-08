package main
import (
    "fmt"
)

func add( a int, b int ) int {
    return a + b
}

func divMod(a int, b int) (int, int) {
    return a / b, a % b
}
func main() {

    fmt.Printf("Simple Addition function : %d ", add(5, 3))
    fmt.Println()
    div, mod := divMod(5, 3)
    fmt.Printf("Return 2 values: Division : %d, Mod : %d", div, mod )
    fmt.Println()
}
