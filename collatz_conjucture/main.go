package main

import (
    "fmt"
)

func collatzSequence(n int) []int {
    seq := []int{n}
    for n != 1 {
        if n != 0 {
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
        fmt.Println(n)
        seq = append(seq, n)
    } else {
        fmt.Println("Reached 0 ")
        return seq
    }
    }
    return seq
}

func main() {
    var n int
    fmt.Print("Enter a positive integer: ")
    _, err := fmt.Scan(&n)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }
    if n <= 0 {
        fmt.Println("Please enter a positive integer.")
        return
    }
    seq := collatzSequence(n)
    fmt.Printf("3x + 1 sequence for %d:\n", n)
    fmt.Println(seq)
}
