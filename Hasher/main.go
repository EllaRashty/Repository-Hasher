package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var client *http.Client

func Hashing(str string) string {
	data := []byte(str)
	hash := base64.StdEncoding.EncodeToString(data)
	return hash
}

func getSha() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var str string
		resp, _ := client.Get("http://host.docker.internal:9091/hash-path")
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&str)
		// fmt.Fprintf(w, "str: %s\n", str)
		hash := Hashing(str)
		fmt.Fprintf(w, "Hash: %s", hash)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Port Repository - 9092:8081\n")
}

func handleRequests() {
	client = &http.Client{Timeout: 10 * time.Second}
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/hashing-service", getSha()).Methods("GET")
	myRouter.HandleFunc("/", start)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
