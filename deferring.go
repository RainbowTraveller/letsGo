package main
/*
defer enables closing the resource only when it is completely used and cleaned up
all necessary things
*/
import (
    "fmt"
)

func cleaning(s string) {
    fmt.Printf("Cleaning resource %s\n", s)
}

func worker() {
    defer cleaning("Resource A")
    defer cleaning("Resource B")
    fmt.Println("Work within the worker\n")
}
func main() {
    worker()
}
