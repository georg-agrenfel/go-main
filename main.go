package main

import "fmt"

func main() {
    var x int = 5
    var y int = 7
    var sum int = add(x, y)
    fmt.Printf("The sum of %d and %d is %d\n", x, y, sum)
}

func add(a int, b int) int {
    return a + b
}
