package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func Scale(v *Vertex,f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {

    f := MyFloat(-5)
    fmt.Printf("%T\n",f)
    fmt.Println(f,f.Abs())

    v := Vertex{3,4}
    fmt.Println(v.Abs())
    fmt.Println(Abs(v))

    v.Scale(10)
    fmt.Println(v)

    Scale(&v,10) // it must be a pointer to structure of type Vertex
    fmt.Println(v)

    p := &Vertex{4,3}
    fmt.Printf("Before scaling: %+v, Abs: %v\n",p,p.Abs())
    p.Scale(10)
    fmt.Printf("After scaling: %+v, Abs: %v\n", p, p.Abs())
    Scale(p,10)
    fmt.Println(p)

}