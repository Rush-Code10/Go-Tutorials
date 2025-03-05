package main

import (
    "fmt"
    "time"
)

// TIME package works on UK Time

func main() {
	
	fmt.Println(time.Now(), "Time now")
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
/*
package main

import (
	"fmt"
)

func main() {
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday":
		fmt.Println("Today is Thursday")
	case "Friday":
		fmt.Println("Today is Friday")
	case "Saturday":
		fmt.Println("Today is Saturday")
	case "Sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid day")
	}
}
*/