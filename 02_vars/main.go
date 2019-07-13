package main

import (
    "fmt"
    "math"
)

var someName = "John" // Global variable

func main() {
    // MAIN TYPES
    // string
    // bool
    // int
    // int int8 int16 int32 int64
    // uint uint8 uint16 uint32 uint64
    // byte - alias for uint8
    // rune - alias for uint32
    // float32 float64
    // complex64 complex128

    // using var (casting as string is optional)
    // there is type inference involved
    var name string = "Kostas" // local variable
    // cannot declare a variable you do not use!
    // compiler does not allow it
    // var lastName = "Gkoutis"

    var age int32 = 37
    otherAge := 34  // This short-hand declaration can only happen when inside of a function
    username, useremail := "Kostas", "Kostas@email.com"
    const isCool = true
    fmt.Println(name, age, isCool, someName, otherAge, username, useremail)
    fmt.Printf("%T\n",name)

    // math
    fmt.Println(math.Floor(5.7)) // 5
    fmt.Println(math.Ceil(2.7)) // 3
    fmt.Println(math.Sqrt(16)) // 4
}
