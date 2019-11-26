package main

import "fmt"

type Skills []string

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	Skills
	int
	class string
}

func main() {
	student1 := Student{Human: Human{"Andi", 34}, class: "Intro to Programming"}

	fmt.Println("Nama saya adalah ", student1.name)
	fmt.Println("Umur saya adalah ", student1.age)
	fmt.Println("Saya ikut kelas ", student1.class)

	student1.Skills = []string{"HTML"}
	fmt.Println("Kemampuan bahasa pemrograman saya adalah ", student1.Skills)
	student1.Skills = append(student1.Skills, "JavaScript", "CSS")
	fmt.Println("Kemampuan bahasa pemrograman saya sekarang adalah ", student1.Skills)

	student1.int = 2
	fmt.Printf("Saya sudah menjadi programmer selama %d tahun\n", student1.int)
}
