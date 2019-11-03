package main

import (
    "fmt"
)

func main() {
    var a [2]string

    a[0] = "Hello"
    a[1] = "world"

    fmt.Println(a[0],a[1])
    fmt.Println(a)

    primes := [6]int{1,2,3,4,5,6}
    fmt.Println(primes)

    names := [4]string{
        "John",
        "Paul",
        "George",
        "Louis",
    }

    fmt.Println(names)

    part1 := names[0:2] // slice of 2 elements from array names
    part2 := names[1:3]

    fmt.Printf("part1 is %T\npart2 is %T\n",part1,part2)

    fmt.Println(part1,part2)

    part2[0] = "XXX"
    fmt.Println(part1,part2)
    fmt.Println(names)

    q := []int{1,2,3,4,5,6}
    fmt.Println(q)

    r := []bool{true,false, true, true, false}
    fmt.Println(r)

    s := []struct { // slice from array of structures
        i int
        b bool
    }{
        {2,true},
        {3,false},
        {4,true},
        {7,false},
        {11,false},
    }

    fmt.Println(s[3])

    slice := []int{2,3,5,7,9,11}

    slice = slice[1:4]
    printSlice("slice",slice)

    slice = slice[:2]
    printSlice("slice",slice)

    slice = slice[1:]
    printSlice("slice",slice)

    // slice-len-cap

    sl := []int{2,3,5,7,11,13}
    printSlice("sl",sl)

    sl = sl[:0]
    printSlice("sl",sl)

    sl = sl[:4]
    printSlice("sl",sl)

    sl = sl[2:] // drop first two values
    printSlice("sl",sl)

    sl = sl[:4]
    printSlice("sl",sl)

    // nil-slices
    var ns []int
    printSlice("ns",ns)
    if ns == nil {
        fmt.Println("nil!")
    }

    // making slices make([]type,len,cap)
    d := make([]int,5)
    printSlice("d",d)

    e := make([]int,0,5)
    printSlice("e",e)

    f := d[:2]
    printSlice("f",f)

    g := f[2:5]
    printSlice("g",g)
}

func printSlice(name string, s []int) {
    fmt.Printf("%s: len=%d cap=%d %v\n",name,len(s),cap(s),s)
}