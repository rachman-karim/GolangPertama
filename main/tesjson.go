package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Air    int    `json:"air"`
	Angin  int    `json:"angin"`
	Wair   string `json:"wair"`
	Wangin string `json:"wangin"`
}

func status(w http.ResponseWriter, r *http.Request) {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err1 := template.ParseFiles("raka.gtpl")
	check(err1)

	rand.Seed(time.Now().UnixNano())
	watermin := 1
	watermax := 10
	waterstat := (rand.Intn(watermax-watermin) + watermin)
	//fmt.Println(waterstat)

	var stwater string
	stwater = ""
	if x := waterstat; x <= 5 {
		stwater = "AMAN"
	} else if x >= 6 && x <= 8 {
		stwater = "SIAGA"
	} else if x > 8 {
		stwater = "BAHAYA"
	}

	windmin := 1
	windmax := 20
	windstat := (rand.Intn(windmax-windmin) + windmin)
	//fmt.Println(windstat)

	var stwind string
	stwind = ""
	if y := windstat; y <= 6 {
		stwind = "AMAN"
	} else if y >= 7 && y <= 15 {
		stwind = "SIAGA"
	} else if y > 15 {
		stwind = "BAHAYA"
	}

	//stwater := "bahaya"
	var object = []Status{{waterstat, windstat, stwater, stwind}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//var jsonString = string(jsonData)
	//fmt.Println(jsonString)
	var data []Status
	var err2 = json.Unmarshal(jsonData, &data)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	fmt.Println(data[0].Angin, data[0].Air, data[0].Wangin, data[0].Wair)

	data2 := struct {
		Title   string
		StAngin string
		StAir   string
		NlAngin int
		NlAir   int
	}{
		Title:   "Siaga Bencana",
		StAngin: data[0].Wangin,
		StAir:   data[0].Wair,
		NlAngin: data[0].Angin,
		NlAir:   data[0].Air,
	}

	err = t.Execute(w, data2)
	check(err)

}

func main() {
	http.HandleFunc("/", status)
	http.HandleFunc("/status", status)       // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}

}
