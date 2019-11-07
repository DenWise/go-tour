package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]byte {
    slice := make([][]byte, dy)

    for i := range slice {
        slice[i] = make([]byte,dx)
        for j := 0; j < dx; j++ {
            slice[i][j] = byte(i * j)
        }
    }

    return slice
}

func main() {
    pic.Show(Pic)
}