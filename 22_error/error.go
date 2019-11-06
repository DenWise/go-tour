package main

import (
    "fmt"
    "time"
    "math"
)

type MyError struct {
    When time.Time
    What string
}

type ErrNegativeSqrt float64

type ErrZeroSqrt float64

func (e ErrZeroSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt zero number: 0")
}

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s\n", e.When, e.What) // returns a string describing error
}

func run() error { // returns a pointer to struct of type "MyError"
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func Sqrt(x float64) (float64,error) {
    if x < 0 {
        return x,ErrNegativeSqrt(x)
    }
    if x == 0 {
        return x,ErrZeroSqrt(x)
    }
	for i, z, prevZ := 1, x / 2, 0.0; ; i++ {
        if prevZ, z = z, z - (z * z - x) / 2 / z; math.Abs(prevZ - z) < 1e-5 { // 1e-5 = 0,00001
            fmt.Printf("Ran %v iterations\n", i)
            return z,nil
        }
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
    fmt.Println(Sqrt(-2))
    fmt.Println(Sqrt(4))
    fmt.Println(Sqrt(0))
}