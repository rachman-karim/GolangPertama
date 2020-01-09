package main

import "fmt"

type Human struct {
    name  string
    age   int
    phone string
}

type Student struct {
    Human  // anonymous field
    school string
}

type Employee struct {
    Human   // anonymous field
    company string
}

// define a method in Human
func (h *Human) SayHi() {
    fmt.Printf("Hi, nama saya %s, umur saya %d dan kalian dapat menghubungi saya di no %s\n", h.name, h.age, h.phone)
}

func main() {
    upin := Employee{Human{"Upin", 22, "123-456-789"}, "Google Inc"}
    ipin := Student{Human{"Ipin", 22, "987-654-321"}, "MIT"}

    upin.SayHi()
    ipin.SayHi()
}
