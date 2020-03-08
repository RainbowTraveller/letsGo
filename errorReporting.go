package main
import (
    "fmt"
    "math"
)

func getsqrt(num float64) (float64, error) {
    if num < 0 {
        return 0.0, fmt.Errorf("Square root of -ve no. is not possible : (%f)", num)
    }
    return math.Sqrt(num), nil
}

func main() {

    num := 25.0
    valid, err := getsqrt(num)

    if err != nil {
        fmt.Printf("Error calculating the square root : %s\n", err)
    } else {
        fmt.Printf("Square root of %f is %f\n", num, valid)
    }
    num = -25.0
    valid, err = getsqrt(num)

    if err != nil {
        fmt.Printf("Error calculating the square root : %s\n", err)
    } else {
        fmt.Printf("Square root of %f is %f\n", num, valid)
    }
}
