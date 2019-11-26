package main

import (
	"errors"
	"fmt"
)

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
}
