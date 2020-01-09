package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's Works!!") // send data to client side
}

func main() {
	http.HandleFunc("/", index)              // set router
	err := http.ListenAndServe(":9090", nil) // set listen port to 9090
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
