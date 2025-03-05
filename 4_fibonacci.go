package main

import "fmt"

func main() {

// FIBONACCI

f := 0
s := 1
fmt.Print(f, " ")
fmt.Print(s, " ")
for i := 0; i < 8; i++ {
t := f + s
fmt.Print(t, " ")
f = s
s = t
}
fmt.Println(" ")
}
