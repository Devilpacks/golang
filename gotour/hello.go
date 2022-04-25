package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
    fmt.Println("Hello, World! \n")
    fmt.Println(math.Pi)
    fmt.Println(add(42, 13))
    a, b := swap("hello", "world")
	fmt.Println(a, b)
    fmt.Println(split(17))
    var i, j int = 1, 2
	fmt.Println(i, c, python, java)
}

