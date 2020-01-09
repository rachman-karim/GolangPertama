package main

import "encoding/json"
import "fmt"

type User struct {
	Status string `json:"Status"`
	Air    int
	Angin  int
}

func main() {
	var object = []User{{"Status", 27, 23}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
