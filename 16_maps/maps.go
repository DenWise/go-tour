package main

import (
    "fmt"
    "golang.org/x/tour/wc"
    "strings"
)

type Vertex struct {
    Lat, Long float64
}

var s = map[string]Vertex{
    "Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func WordCount(s string) map[string]int {
    words := make(map[string]int)

    for _, w := range strings.Fields(s) {
        if elem, ok := words[w]; ok {
            words[w] = elem + 1
        } else {
            words[w] = 1
        }
    }

    return words
}

func main() {
    m := make(map[string]Vertex) // making is necessary
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
    fmt.Println(s)

    key := "Google"

    delete(s,key)
    l,ok := s[key]
    fmt.Printf("Location %v is present? %v\n",l,ok)

    wc.Test(WordCount)
}