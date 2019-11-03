package main

import (
    "fmt"
    "strings"
)

func main() {
    board := [][]string{
        []string{"_","_","_"},
        []string{"_","_","_"},
        []string{"_","_","_"},
    }

    board[0][0] = "X"
    board[2][2] = "O"
    board[1][2] = "X"
    board[1][0] = "O"
    board[0][2] = "X"

    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }

    var s []int
    printSlice("s",s)

    s = append(s,0)
    printSlice("s",s)

    s = append(s,1)
    printSlice("s",s)

    s = append(s,2,3) // cap = 4
    printSlice("s",s)

    s = append(s,4,5) // cap = 8 ?why?
    printSlice("s",s)

    var pow = []int{1,2,4,8,16,32,64,128}

    for i,v := range pow {
        fmt.Printf("2**%d = %d\n",i,v)
    }

    poww := make([]int,10)

    for i := range poww {
        poww[i] = 1 << uint(i)
    }
    
    for _,v := range poww {
        fmt.Printf("%d\n",v)
    }
}

func printSlice(name string, s []int) {
    fmt.Printf("%s: len=%d cap=%d %v\n",name,len(s),cap(s),s)
}