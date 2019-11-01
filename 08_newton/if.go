package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	for i, z, old := 1, x / 2, 0.0; ; i++ {
		fmt.Println(old,z)
        if old, z = z, z - (z * z - x) / 2 / z; math.Abs(old - z) < 1e-6 {
            fmt.Printf("Ran %v iterations\n", i)
            return z
        }
    }
}
 
func main() {
	fmt.Println(Sqrt(3),math.Sqrt(3))
}