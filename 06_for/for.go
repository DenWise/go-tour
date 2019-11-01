package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum2 := 1
	for sum2 < 10 { // like while
		sum2 += sum2
	}
	fmt.Println(sum2)

	// for {} - infinit loop, never do it
}