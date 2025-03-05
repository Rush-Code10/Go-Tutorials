package main

import (
"fmt"
"math"
)

const s string = "constant"

func main() {
fmt.Println(s)

const n = 500000000

const d = 3e20 / n // 3e20 --> 3*10(base)^20 --> 300000000000000000000

fmt.Println(d)

fmt.Println(int64(d))

fmt.Println(math.Sin(n))
fmt.Println(math.Cos(n))
}