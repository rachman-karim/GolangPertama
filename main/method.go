package main

import (
    "fmt"
    "math"
)

type Circle struct {
    radius float64
}

type Rectangle struct {
    width, height float64
}

// method
func (c Circle) Area() float64 {
    return c.radius * c.radius * math.Pi
}

// method
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    c1 := Circle{10}
    c2 := Circle{25}
    r1 := Rectangle{9, 4}
    r2 := Rectangle{12, 2}
    fmt.Println("Area Circle 1: ", c1.Area())
    fmt.Println("Area Circle 2: ", c2.Area())
    fmt.Println("Area Rectangle 1: ", r1.Area())
    fmt.Println("Area Rectangle 2: ", r2.Area())
}
