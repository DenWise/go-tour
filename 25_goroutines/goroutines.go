package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
  go say("Hello1")
  say("Hello2")

  s := []int{7,2,8,-9,4,0,6,7,8,9,5,0,7,2,8,-9,4,0,6,7,8,9,5,0}

  c := make(chan int, 3) // buffering channel 3 - buffer size / cap(c)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  x, y := <-c, <-c // receive from c

  fmt.Println(x, y, x+y)

  channel := make(chan int, 10)
  go fibonacci(cap(channel),channel)
//   for i := range channel {
//       fmt.Println(i)
//   }

  for {
      v, ok := <- channel
      if !ok {
          break
      }
      fmt.Println(v)
  }
}