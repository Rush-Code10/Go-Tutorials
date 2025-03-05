package main

import "fmt"

func main() {
fmt.Println("Hello World")
fmt.Println("7.5/3.5 = ", 7.5/3.5)
fmt.Println(true || false)
fmt.Println(!true)

var a = "hey" // Declaring a String variable
var b, c int = 1, 2 // Declaring Integer variables
d := 10 // Other way to declare (without datatypes)
var e int // Empty declaration
f := 100 // Other way to declare (without datatypes)
_ = f // Comment this out, will throw an error

// In GO every variable has to be used somewhere if its not, it will throw an error

fmt.Println(a, b, c, d, e)
}