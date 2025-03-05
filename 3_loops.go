package main

import "fmt"

func main() {

// 1
fmt.Println("Method One")
i := 1
for i <= 3 {
fmt.Println(i)
i = i + 1
}

// 2
fmt.Println("Method Two")
for j := 0; j < 3; j++ {
fmt.Println(j)
}

// 3
fmt.Println("Method Three")
for k := range 3 {
fmt.Println("range", k)
}

// 4
fmt.Println("Using break")
for {
fmt.Println("LOOOP")
break
}

// 5
fmt.Println("Using continue")
for n := range 7 {
if n%2 != 0 {
continue
}
fmt.Println(n)
}
}