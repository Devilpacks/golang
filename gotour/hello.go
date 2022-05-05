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
	fmt.Println(i, j, c, python, java)
}


// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
func shortAssignment() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "nooooooooooooooooo!"

	fmt.Println(i, j, k, c, python, java)
}