package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

type F float64

func (f F) M() {
    fmt.Println(f)
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f // a MyFLoat implements Abser
    a = &v // a *Vertex implements Abser

    fmt.Println(a.Abs())

    var i I = &T{"Hello"}
    i.M()

    var inf I

    inf = &T{"Hello"}
    describe(inf)
    inf.M()

    inf = F(math.Pi)
    describe(inf)
    inf.M()

    var t *T
    inf = t
    describe(inf) // nil, *main.T
    inf.M() // nil

    inf = &T{"Hello"} // pointer to struct of type T
    describe(inf)
    inf.M()

    var voidI interface {}
    describeVoid(voidI)

    voidI = 42
    describeVoid(voidI)

    voidI = "Gopher"
    describeVoid(voidI)

    // type-assertion
    var ii interface {} = "hello"

    s := ii.(string) // returns a value from interface
    fmt.Println(s)

    s, ok := ii.(string)
    fmt.Println(s,ok) // s = value, ok = true/false

    ff, ok := ii.(float64)
    fmt.Println(ff, ok)

    ff = ii.(float64)
    fmt.Println(ff)
}

func describe(i I) {
    fmt.Printf("(%v,%T)\n",i,i)
}

func describeVoid(i interface{}) {
    fmt.Printf("(%v,%T)\n",i,i)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}