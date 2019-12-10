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

// return greater value between parameters
func max(param1, param2 int) int {
	if param1 > param2 {
		return param1
	}
	return param2
}

func sumAndTime(param1, param2 int) (int, int) {
	return param1 + param2, param1 * param2
}

type testInt func(int) bool // define a function type of variable
var data = []int{1, 2, 3, 4, 5, 6, 7}

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
	arrayTest := []int{74, 13, 22, 49}
	for index, value := range arrayTest {
		fmt.Printf("Index %d = %d\n", index, value)
	}

	slice := make([]int, 0, 10)
	slice = append(slice, 8)
	fmt.Println(slice) // [8]

	slice2 := make([]int, 10)
	slice2 = append(slice2, 8)
	//slice2 = slice2[0:8]
	//slice2[7] = 8
	fmt.Println(slice2)

	scores := []int{1, 2, 3, 4, 5}
	slice4 := scores[2:4]
	slice4[0] = 999

	fmt.Println("slice: ", slice4)
	fmt.Println("scores: ", scores)

	scores2 := []int{1, 2, 3, 4, 5}
	newScores := make([]int, 5)

	copy(newScores, scores2)
	newScores[0] = 999
	scores2[1] = 1000

	fmt.Println("copy: ", newScores)
	fmt.Println("scores: ", scores2)

	//MAP
	var numbers2 map[string]int
	// atau
	numbers2 = make(map[string]int)
	numbers2["one"] = 1 // assign value by key
	numbers2["ten"] = 10
	numbers2["six"] = 6

	fmt.Println(numbers2["ten"]) // get value

	// initialize map with key and value
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	// map memiliki 2 return value, jika key tersedia pada map maka 'exist' akan true tetapi bila tidak ada maka akan false
	cSharpRating, exist := rating["C#"]
	if exist {
		fmt.Println("C# tersedia pada map dan memiliki value ", cSharpRating)
	} else {
		fmt.Println("C# tidak tersedia pada map")
	}

	// delete elemen
	delete(rating, "C")
	//-------------------------------------------------------
	no1 := 7
	no2 := 2

	fmt.Printf("max(%d, %d) = %d\n", no1, no2, max(no1, no2))
	fmt.Printf("max(%d, %d) = %d\n", no2, no1, max(no2, no1))

	//-------------------------------------------------------
	no11 := 7
	no21 := 2

	no1PLUSno2, no1TIMESno2 := sumAndTime(no11, no21)

	fmt.Printf("%d + %d = %d\n", no1, no2, no1PLUSno2)
	fmt.Printf("%d * %d = %d\n", no2, no1, no1TIMESno2)

	//-------------------------------------------------------
	variadicFunc(4, 11, 23, 7, 33, 10, 5)

	//-------------------------------------------------------
	fmt.Printf("maxx(%d, %d) = %d\n", no1, no2, maxx(no1, no2))
	fmt.Printf("maxx(%d, %d) = %d\n", no2, no1, maxx(no2, no1))

	//-------------------------------------------------------
	odd := filter(data, isOdd)
	even := filter(data, isEven)

	fmt.Println("Data = ", data)
	fmt.Println("Bilangan ganjil = ", odd)
	fmt.Println("Bilangan genap = ", even)

	//---------------------------------------------------------
	numbers := make(map[string]int)
	numbers["one"] = 1 // assign value by key
	numbers["ten"] = 10
	numbers["six"] = 6

	fmt.Println(numbers["three"]) // get value

	//--------------------------------------------------------
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

}

func variadicFunc(args ...int) {
	for key, val := range args {
		fmt.Printf("Parameter index ke %d = %d\n", key, val)
	}
}

// return greater value between parameters
func maxx(param1, param2 int) int {
	if param1 > param2 {
		return param1
	}
	return param2
}

func isOdd(integer int) bool {
	return integer%2 != 0
}

func isEven(integer int) bool {
	return integer%2 == 0
}

func filter(dt []int, fn testInt) []int {
	var result []int
	for _, value := range dt {
		if fn(value) {
			result = append(result, value)
		}
	}
	return result
}
