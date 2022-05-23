package main

import (
	"fmt"
	"math/cmplx"
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

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func zeroValues() {
	// 	0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConversions() {
	// The expression T(v) converts the value v to the type T.
	// i := 42
	// f := float64(i)
	// u := uint(f)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func main() {
    fmt.Println("Hello, World! \n")
    fmt.Println(math.Pi)
    fmt.Println(add(42, 13))
    a, b := swap("hello", "world")
	fmt.Println(a, b)
    fmt.Println(split(17))
    var i, j int = 1, 2
	fmt.Println(i, j, c, python, java)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	zeroValues()
	shortAssignment()
	typeConversions()
}


// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
func shortAssignment() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "nooooooooooooooooo!"
	fmt.Println(i, j, k, c, python, java)
}


