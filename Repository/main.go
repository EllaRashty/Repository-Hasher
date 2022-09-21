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
	Name     string `json:"name"`
	Path     string `json:"path"`
	Size     int    `json:"size"`
	Sha      string `json:"sha"`
	Url      string `json:"url"`
	Type     string `json:"type"`
	Html     string `json:"html_url"`
	Download string `json:"download_url"`
}

type ContentsFacts []ContentFact

var ref = "master"
var currentPath string
var strToHash = ""

func GetFileContents(pathInRepo string) ContentsFacts {
	var contentsFacts ContentsFacts
	err := GetJson("https://api.github.com/repos/"+pathInRepo+"/contents/?ref="+ref, &contentsFacts)
	if err != nil {
		fmt.Printf("error getting contents facts: %s\n", err.Error())
	}
	fmt.Printf("contents of: %s\n", pathInRepo)
	return contentsFacts
}

func getPathsForHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var path string
		if err := json.NewDecoder(r.Body).Decode(&path); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		parts := strings.SplitAfter(path, ",")
		// arr := strings.FieldsFunc(path, func(r rune) bool {
		// 	return r == ','
		// })
		for _, x := range parts {
			fmt.Fprintf(w, "%s\n", x)
		}
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
		for _, pathInRepo := range pathInRepos {
			fmt.Printf("yey? :%s", pathInRepo)
			allSha = allSha + collectSha(pathInRepo)
		}
	}
	strToHash = allSha
}

func collectSha(pathInRepo string) string {
	contentsFacts := GetFileContents(pathInRepo)
	str := ""
	for _, x := range contentsFacts {
		str = str + x.Sha
	}
	return str
}

func CheckoutRef() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i string
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ref = i
		GetFileContents(currentPath)
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

func setPath() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&currentPath); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		contentsFacts := GetFileContents(currentPath)
		for i, x := range contentsFacts {
			fmt.Fprintf(w, "%d.\nName: %s,\nPath: %s,\nSha: %s,\nSize: %d,\nUrl: %s,\nType: %s,\nHtml: %s,\nDownload: %s,\n\n", i+1, x.Name, x.Path, x.Sha, x.Size,
				x.Url, x.Type, x.Html, x.Download)
		}
	}
}

func getCurrentPath() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(currentPath); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getHash() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(strToHash); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Port Repository - 9091:8081\n")
}

func handleRequests() {
	client = &http.Client{Timeout: 10 * time.Second}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/input-path", setPath()).Methods("POST")
	myRouter.HandleFunc("/checkout", CheckoutRef()).Methods("POST")
	myRouter.HandleFunc("/input-path", getCurrentPath()).Methods("GET")
	myRouter.HandleFunc("/hash-path", getHash()).Methods("GET")
	myRouter.HandleFunc("/hash-path", getPathsForHash()).Methods("POST")
	myRouter.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
