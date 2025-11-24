package main

import "net/http"
import "encoding/json"

/* 	initial Todo struct, representing a single item on a to-do list. 
	Could be expanded with additional fields, like times/dates, priorities, etc.) if wanted.	*/
type Todo struct {

	// using JSON "struct tags" so that when the struct is converted to/from JSON the field names are kept.
    Title       string `json:"title"`
    Description string `json:"description"`
}

// basic array to hold all Todos in memory
var todos []Todo

func main() {
	// Your code here
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Your code here
}
