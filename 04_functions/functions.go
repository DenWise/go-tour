package main

import (
	"fmt"
)
var c, python, java = true, false, "no!"

func add(x,y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x,y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42,13))
	a, b := swap("lang","Go")
	fmt.Println(a,b)
	x, y := split(add(42,13))
	fmt.Println(x, y)
	fmt.Println(c, python, java)
}