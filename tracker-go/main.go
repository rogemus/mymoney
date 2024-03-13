package main

import (
	"fmt"
	"log"
	"net/http"
)

const port_num string = ":8877"

// Handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

// Info
func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info")
}

func main() {
	log.Println("Starting http server.")

	// Register handlers
	http.HandleFunc("/", Home)
	http.HandleFunc("/info", Info)

	log.Println("Stared on port: ", port_num)
	log.Println("To close connection press CTRL+C")

	err := http.ListenAndServe(port_num, nil)
	if err != nil {
		log.Fatal(err)
	}
}
