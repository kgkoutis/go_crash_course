package main

import (
    "fmt"
)

func main() {
    // Arrays
    var fruitArr [2] string

    // Assign values
    fruitArr [0] = "Apple"
    fruitArr [1] = "Orange"

    fmt.Println(fruitArr) // [Apple Orange]
    fmt.Println(fruitArr[0]) // Apple

    // Declare and assign
    fruitArr2 := [2]string{"Apple", "Banana"}
    // fruitArr3 := [2]{"Apple", "Banana"} // shorthand notation does not work here!
    fmt.Println(fruitArr2) // [Apple Banana]

    // fruitNotASlice := [2]string{"Apple", "Orange", "Grape", "Cherry"} // This would fail, more elements in the array
    fruitSlice := []string{"Apple","Orange","Grape", "Cherry"} // You define a slice by not specifying the number
    // fruitSlice2 := []{"Apple","Orange","Grape", "Cherry"} // Shorthand notation does not work here!
    fmt.Println(fruitSlice) // [Apple Orange Grape Cherry]
    fmt.Println(len(fruitSlice)) // 4
    fmt.Println(fruitSlice[1:3]) // [Orange Grape]
    fmt.Println(fruitSlice[1:]) // [Orange Grape Cherry]
}
