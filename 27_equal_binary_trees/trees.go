package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

func walkTree(t *tree.Tree,ch chan int) {
    if t.Left != nil {
        walkTree(t.Left,ch)
    }
    ch <- t.Value
    if t.Right != nil {
        walkTree(t.Right,ch)
    }
}

func Walk(t *tree.Tree, ch chan int) {
        walkTree(t,ch)
        close(ch)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go Walk(t1,ch1)
    go Walk(t2,ch2)
    for v1 := range ch1 {
        v2,ok := <- ch2
        if !ok || v1 != v2 {
            return false
        }
    }
    return true
}

func main() {
    c := make(chan int)
    t := tree.New(1)
    go Walk(t,c);
    for v := range c {
        fmt.Println(v)
    }

    fmt.Printf("Must be true: %v\n", Same(tree.New(1),tree.New(1)))
    fmt.Printf("Must be false: %v\n", Same(tree.New(3),tree.New(4)))
}