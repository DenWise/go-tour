package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3,4)
}

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func fibonacci() func() int {
    prev, cur, cnt := 0, 1, 0
    return func() int {
        if cnt == 0 {
            cnt++
            return prev
        }
        if cnt == 1 {
            cnt++
            return cur
        }
        next := prev + cur
        prev, cur = cur, next
        cnt++
        return next
    }
}

func main() {
    hypot := func(x,y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println("hypot(5,12)",hypot(5,12))

    fmt.Println("compute(hypot)",compute(hypot))
    fmt.Println("3^4",compute(math.Pow))

    // closures

    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
    fmt.Println(pos(10))

    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}