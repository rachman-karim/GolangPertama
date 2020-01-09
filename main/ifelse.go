package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	watermin := 1
	watermax := 10
	waterstat := (rand.Intn(watermax-watermin) + watermin)
	var stwater string
	stwater = ""
	if x := waterstat; x < 6 {
		stwater = "aman"
	} else if x >= 6 && x <= 8 {
		stwater = "siaga"
	} else if x > 8 {
		stwater = "bahaya"

	}
	fmt.Println(stwater, waterstat)
}
