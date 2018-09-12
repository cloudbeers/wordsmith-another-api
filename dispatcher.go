package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/version", whichVersion)

	fmt.Println("Listening on port 80")
	http.ListenAndServe(":80", nil)
}

func whichVersion(w http.ResponseWriter, r *http.Request) {

	version, err := ioutil.ReadFile("/VERSION")
	if err != nil {
		log.Fatal(err)
	}

	currentVersion := string(version)
	fmt.Printf("Running version: %s", currentVersion)

	js := fmt.Sprintf("{ \"version\": \"%s\" }", currentVersion)

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, js)
}
