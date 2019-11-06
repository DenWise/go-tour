package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int,error) {
    for x := range b {
        b[x] = 'A'
    }
    return len(b),nil
}

type rot13Reader struct {
    r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int,error) {
    rot.r.Read(b)
    for i := 0; i < len(b); i++ {
        if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
            b[i] += 13
        } else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
            b[i] -= 13
        }
    }
    return len(b),io.EOF
}

func main() {
    r := strings.NewReader("Hello, Reader!")

    b := make([]byte,5)
    for {
        n, err := r.Read(b)
        fmt.Printf("n = %v; err = %v; b = %v\n",n,err,b)
        fmt.Printf("b[:n] = %q\n",b[:n])
        if err == io.EOF {
            break
        }
    }

    reader.Validate(MyReader{})

    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    rot := rot13Reader{s}
    io.Copy(os.Stdout, &rot)
}