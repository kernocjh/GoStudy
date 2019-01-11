package main

import (
    "fmt"
    "gopl/ch1/calc/math"
)

func main() {
    var x int = 5
    var y int = 7
    fmt.Printf("%d + %d = %d\n", x, y, math.Add(x, y))
    fmt.Printf("%d - %d = %d\n", x, y, math.Sub(x, y))
    fmt.Printf("%d %% %d = %d\n", x, y, math.Mod(x, y))
}
