package main

import (
    "fmt"
)

type Vertex struct {
    X, Y int
}

func main() {
    fmt.Println(Vertex{1,2})
    v := Vertex{2,5}
    v.X = 10
    fmt.Println(v.X)

    p := &v
    p.X = 1e9
    fmt.Println(v)

    v1 := Vertex{X:3}
    v2 := Vertex{1,4}
    v3 := Vertex{}
    pointer := &Vertex{1,4}

    fmt.Println(v1,v2,pointer.Y,v3)
}