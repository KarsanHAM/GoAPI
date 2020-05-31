package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Courses []Course

func allCourses(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	courses := Courses{
		Course{Title: "Test Title", Desc: "Description", Content: "test123"},
	}

	fmt.Println("Endpoint hit: All articles endpoint")
	json.NewEncoder(w).Encode(courses)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Testpost endpoint worked!")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/courses", allCourses).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	handleRequests()
}
