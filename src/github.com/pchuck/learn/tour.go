package main

import (
	"fmt"
	"math"
	"math/cmplx"
)


// a simple function
func multiply(x, y int) int {
	return x * y
}

// functions return multiple parameters
func swap(x, y string) (string, string) {
	return y, x
}

// functions return named parameters
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// variables can be declared in a list, type at end
func printVars() {
	var (
		i int = 1 // with explicit type
		c, python, java = true, false, "yes" // implicit type
	)
	k := 3 // short assignment, available in functions
	var z complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Println(i, c, python, java, k, z)
}

func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)
}

func needInt(x int) int { return x * 10 + 1 }
func constants() {
	const World = "世界"
	const (
		Big = 1 << 100
		Small = Big >> 99
	)
	fmt.Println(World)
//	fmt.Println(needInt(Big))
	fmt.Println(Small)
	
}

func forLoop () {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println(sum)

	s := 1
	for s < 10 {
		s += s
		fmt.Print(s)
		fmt.Print(" ")
	}

	fmt.Println()
}

func main() {
	fmt.Println(multiply(42, 13))
	fmt.Println(swap("a", "b"))
	fmt.Println(split(17))
	printVars()
	typeConversion()
	constants()
	forLoop()
}
