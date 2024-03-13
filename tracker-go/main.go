package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port_num string = ":8877"

// Helpers
func composeErrorJsonData(msg string) string {
	data := map[string]interface{}{
		"error": msg,
	}

	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}

// Handler
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"hello": "world",
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		errData := composeErrorJsonData("Could not marshal json %s\n")
		fmt.Fprintf(w, errData)
	}

	fmt.Fprintf(w, string(jsonData))
}

// Info
func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Info")
}

func main() {
	log.Println("Starting http server.")

	// Register handlers
	http.HandleFunc("/", Home)

	log.Println("Stared on port: ", port_num)
	log.Println("To close connection press CTRL+C")

	err := http.ListenAndServe(port_num, nil)
	if err != nil {
		log.Fatal(err)
	}
}
