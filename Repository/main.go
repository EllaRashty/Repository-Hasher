package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var client *http.Client

type ContentFact struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Size int    `json:"size"`
	Sha  string `json:"sha"`
	Url  string `json:"url"`
	Type string `json:"type"`
	Html string `json:"html_url"`
}

type ContentsFacts []ContentFact

var ref = "master"
var strToHash = ""

func GetFileContents(pathInRepo string) ContentsFacts {
	var contentsFacts ContentsFacts
	err := GetJson("https://api.github.com/repos/"+pathInRepo+"/contents/?ref="+ref, &contentsFacts)
	if err != nil {
		fmt.Printf("error getting contents facts: %s\n", err.Error())
	}
	fmt.Printf("contents of: %s , ref: %s\n", pathInRepo, ref)
	return contentsFacts
}

func getPathsForHash() http.HandlerFunc { // Get input from user of path in order to send it to the hashing microservice
	return func(w http.ResponseWriter, r *http.Request) {
		var path string
		if err := json.NewDecoder(r.Body).Decode(&path); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		parts := strings.SplitAfter(path, ",") // divide each path

		// check for the input paths
		// for _, x := range parts {
		// 	fmt.Fprintf(w, "--- %s\n", x)
		// }

		HashFiles(path)

		fmt.Fprintf(w, "The Path For Hashing: %s\n", parts)
	}
}

func HashFiles(pathInRepos ...string) {
	allSha := ""
	if len(pathInRepos) == 1 {
		allSha = collectSha(pathInRepos[0])
	} else {
		// make sure the hashing will run in the same order no matter to the inputs order
		sort.Strings(pathInRepos)
		for _, pathInRepo := range pathInRepos { // Collect all tha SHA from each path
			fmt.Printf("path :%s \n", pathInRepo)
			allSha = allSha + collectSha(pathInRepo)
		}
	}
	strToHash = allSha
}

func collectSha(pathInRepo string) string { //  Collect all tha SHA from each file
	contentsFacts := GetFileContents(pathInRepo)
	allSha := ""
	for _, file := range contentsFacts {
		allSha = allSha + file.Sha
	}
	return allSha
}

func CheckoutRef() http.HandlerFunc { // Get input from user to change to a different ref
	return func(w http.ResponseWriter, r *http.Request) {
		var newRef string
		if err := json.NewDecoder(r.Body).Decode(&newRef); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ref = newRef
		fmt.Fprintf(w, "The Current gitRef is: %s", ref)
	}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func setPath() http.HandlerFunc { // Get input from user of specific repo to display his data
	return func(w http.ResponseWriter, r *http.Request) {
		var currentPath string
		if err := json.NewDecoder(r.Body).Decode(&currentPath); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		contentsFacts := GetFileContents(currentPath)
		for index, file := range contentsFacts {
			fmt.Fprintf(w, "%d.\nName: %s,\nPath: %s,\nSha: %s,\nSize: %d,\nUrl: %s,\nType: %s,\nHtml: %s\n\n", index+1, file.Name, file.Path, file.Sha, file.Size,
				file.Url, file.Type, file.Html)
		}
	}
}

// func getCurrentPath() http.HandlerFunc { // Check for validation
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(currentPath); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }

func getHash() http.HandlerFunc { // Allowing the hashing microservice (other port) to send GET request  (to the Repo microservice)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(strToHash); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) { // Default page
	fmt.Fprintf(w, "Port Repository - 9091:8081\n")
}

func handleRequests() {
	client = &http.Client{Timeout: 10 * time.Second}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/input-path", setPath()).Methods("POST")   // Get input from user of specific repo to display his data
	myRouter.HandleFunc("/checkout", CheckoutRef()).Methods("POST") // Get input from user to change to a different ref
	// myRouter.HandleFunc("/input-path", getCurrentPath()).Methods("GET") // Check for validation
	myRouter.HandleFunc("/hash-path", getHash()).Methods("GET")          // Allowing the hashing microservice (other port) to send GET request  (to the Repo microservice)
	myRouter.HandleFunc("/hash-path", getPathsForHash()).Methods("POST") // Get input from user of path in order to send it to the hashing microservice
	myRouter.HandleFunc("/", homePage)                                   // Default page
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
