package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	// Print het request type
	fmt.Fprintf(w, "Request-type:%s\n", r.Method)

	// Loop over headers
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s \n", name, value)
		}
	}

	// Vanaf hier volgt de Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s\n", body)

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {

	port := getEnv("HTTP_PORT", "8888")

	http.HandleFunc("/", requestHandler)

	fmt.Println("Started, serving at ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
