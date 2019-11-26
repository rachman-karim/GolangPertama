package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	class string
	phone string
}

func main() {
	student1 := Student{Human{"Andi", 34, "+6221-777-444"}, "Intro to Programming", "+6222-987-654"}

	fmt.Println("No telfon rumah saya adalah ", student1.Human.phone)
	fmt.Println("No telfon pribadi saya adalah ", student1.phone)
}
