package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array elements:", arr)

    arr[2] = 10 // Change the 3rd element (index 2) to 10
    fmt.Println("Modified array elements:", arr)
}
