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

var str string

func Hashing(temp string) string {
	data := []byte(temp)
	hash := base64.StdEncoding.EncodeToString(data)
	return hash
}

func putHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var str string
		resp, _ := client.Get("http://localhost:9091/hasher")
		defer resp.Body.Close()
		json.NewDecoder(resp.Body).Decode(&str)
		hash := Hashing(str)
		fmt.Fprintf(w, "amen: %s", hash)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Port Repository - 9092:8081\n")
}

func handleRequests() {
	client = &http.Client{Timeout: 10 * time.Second}
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/hashing", putHash()).Methods("GET")
	myRouter.HandleFunc("/", start)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}

// package main

// import (
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/mux"
// )

// type Str struct {
// 	Name string `json:"name"`
// }

// var client *http.Client

// func Hashing(str string) string {
// 	data := []byte(str)
// 	hash := base64.StdEncoding.EncodeToString(data)
// 	return hash
// }

// func putHash() string {
// 	var str string
// 	resp, _ := client.Get("http://localhost:9091/hasher")
// 	defer resp.Body.Close()
// 	json.NewDecoder(resp.Body).Decode(&str)
// 	return Hashing(str)
// }

// func start(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Port Repository - 9092:8081\n")
// 	// putHash()
// 	fmt.Printf("fact: %s\n", putHash())
// }

// func handleRequests() {
// 	client = &http.Client{Timeout: 10 * time.Second}
// 	myRouter := mux.NewRouter().StrictSlash(true)

// 	myRouter.HandleFunc("/", start)
// 	log.Fatal(http.ListenAndServe(":8081", myRouter))
// }

// func main() {
// 	handleRequests()
// }
