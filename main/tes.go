package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Skills []string
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	Skills
	int
	class string
}

type Murid struct {
	Human
	class string
	phone string
}

type person struct {
	name string
	age  int
}

func testFunc() {
	number := 0

testLabel:
	fmt.Println(number)
	number++
	goto testLabel
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

	student1 := Student{Human: Human{"Andi", 34, ""}, class: "Intro to Programming"}

	fmt.Println("Nama saya adalah ", student1.name)
	fmt.Println("Umur saya adalah ", student1.age)
	fmt.Println("Saya ikut kelas ", student1.class)

	student1.Skills = []string{"HTML"}
	fmt.Println("Kemampuan bahasa pemrograman saya adalah ", student1.Skills)
	student1.Skills = append(student1.Skills, "JavaScript", "CSS")
	fmt.Println("Kemampuan bahasa pemrograman saya sekarang adalah ", student1.Skills)

	student1.int = 2
	fmt.Printf("Saya sudah menjadi programmer selama %d tahun\n", student1.int)

	student2 := Murid{Human{"Andi", 34, "+6221-777-444"}, "Intro to Programming", "+6222-987-654"}

	fmt.Println("No telfon rumah saya adalah ", student2.Human.phone)
	fmt.Println("No telfon pribadi saya adalah ", student2.phone)

	// inisialisasi variabel
	//number := rand.Intn(20);
	number := rand.Intn(20)
	//if number := rand.Intn(20); number < 10 {
	if number < 10 {
		fmt.Println("Nomor", number, "lebih kecil dari 10")
	} else if number == 10 {
		fmt.Println("Nomor", number, "sama dengan 10")
	} else {
		fmt.Println("Nomor", number, "lebih besar dari 10")
	}

	//testFunc()

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	/*
		sum2 := 0
		for sum2 < 10 {
			sum2 += sum2
		}
	*/
	fmt.Println(sum)

	//switch case fallthrough = untuk tidak keluar setelah konsidi terpenuhi

	index := 10
	switch index {
	case 1:
		fmt.Println("index <= 1")
		fallthrough
	case 2, 3, 4:
		fmt.Println("index <= 2 atau 3 atau 4")
		fallthrough
	case 10:
		fmt.Println("index <= 10")
		fallthrough
	default:
		fmt.Println("index adalah bilangan integer")
	}

	//Array
	arrayTest := [4]int{74, 13, 22, 49}
	for index, value := range arrayTest {
		fmt.Printf("Index %d = %d\n", index, value)
	}

}
