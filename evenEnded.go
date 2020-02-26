package main
import (
    "fmt"
)


func main() {
    count := 0
    for i:= 1000; i <= 9999; i=i+1 {
        for j:= i; j <= 9999; j=j+1 {

            num := i * j;
            numStr := fmt.Sprintf("%d", num)

            if numStr[0] == numStr[len(numStr) - 1] {
                //fmt.Printf("Num %d\n : Even ended", num)
                count = count + 1
            } else {
                ///fmt.Printf("Num %d\n : Not Even ended", num)
            }
        }
    }
    fmt.Printf("Count : %d", count)
}
