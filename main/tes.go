package main

import (
	"errors"
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	class string
}
type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Rachman")
	fmt.Println(len("It's works!"))
	fmt.Println("It's works!")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1.2 + 1.3 =", 1.2+1.3)

	err := errors.New("Data not complete!")
	err = nil
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("raka oke")
	}

	var P person

	P.name = "Andi"
	P.age = 34
	//P := struct {		name string		age  int	}{"Andi", 34}
	fmt.Println("Nama saya adalah ", P.name)

	student1 := Student{Human{"Andi", 34}, "Intro to Programming"}

	fmt.Println("Nama saya adalah ", student1.name)
	fmt.Println("Umur saya adalah ", student1.age)
	fmt.Println("Saya ikut kelas ", student1.class)

}
