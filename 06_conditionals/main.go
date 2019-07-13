package main

import (
    "fmt"
)

func main() {
    x := 5
    y :=10

    if x <= 10 { // common practice is not to use parens
        fmt.Printf("%d is less than %d\n", x, y)
    } else {
        fmt.Printf("%d is less than %d\n", y, x)
    }

    color := "red"

    if color == "red" {
        fmt.Println("s")
    } else {
        fmt.Println("color is blue")
    } else {
        fmt.Println("color is NOT blue or red")
    }

    // Switch
    switch color {
    case "red":
        fmt.Println("color is red")
    case "blue":
        fmt.Println("color is blue")
    default:
        fmt.Println("color is NOT blue or red")
    }
}
