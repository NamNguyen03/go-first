package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var persons = []Person{
	{Name: "John", Age: 25},
	{Name: "Alice", Age: 30},
	{Name: "Bob", Age: 35},
	{Name: "Eve", Age: 40},
}

func getAllPersons(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(persons)
}

func greetings(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello go web! with Go"))
}

func main() {
	http.HandleFunc("/", greetings)
	http.HandleFunc("/persons", getAllPersons)
	http.ListenAndServe(":8080", nil)
}
