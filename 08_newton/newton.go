package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	for i, z, prevZ := 1, x / 2, 0.0; ; i++ {
        if prevZ, z = z, z - (z * z - x) / 2 / z; math.Abs(prevZ - z) < 1e-5 { // 1e-5 = 0,00001
            fmt.Printf("Ran %v iterations\n", i)
            return z
        }
    }
}
 
func main() {
	fmt.Println(Sqrt(3),math.Sqrt(3))
}