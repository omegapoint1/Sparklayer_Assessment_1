package main

import "net/http"
import "encoding/json"

/* 	initial Todo struct, representing a single item on a to-do list 
	Could be expanded with additional fields, like times/dates, priorities, etc.) if wanted	*/
type Todo struct {

	// using JSON "struct tags" so that when the struct is converted to/from JSON the field names are kept
    Title       string `json:"title"`
    Description string `json:"description"`
}

// basic array to hold all Todos in memory
var todos []Todo

func main() {
	// when svelte frontend makes a request to the root path (which it already does), call ToDoListHandler
	http.HandleFunc("/", ToDoListHandler)

	// start the server on port 8080. nil should cause root requests to be handled by ToDoListHandler
	http.ListenAndServe(":8080", nil)
}

// w is the "writer", meaning it used to send a response back to the frontend
// r is the "request", meaning it contains all information from the frontends request
func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//using a switch statement instead of multiple if-elses for clarity
	switch r.Method {

	// when svelte frontend sends a GET request, it expects a JSON containing all Todo items
	case "GET":
		// Try to send the whole array of Todo items as JSON, which the frontend expects.

		// create a new JSON encoder using the writer, which should send data back to the frontend correctly
		encoder := json.NewEncoder(w)

		// encode the Todo array as JSON using the encoder object. This seems to send the data back to the frontend directly.
		err := encoder.Encode(todos)

		//the status code should be returned automatically by the encoder (200 on success) as per the spec

		// basic error handling
        if err != nil {
			http.Error(w, "Unable to encode todos as JSON", http.StatusInternalServerError)
			return
		}


	// if svelte frontend sends a POST request, it is submitting a new Todo item to be added to the list
	case "POST":
		//Try to decode the JSON sent by the frontend into a new Todo struct and add it to the memory array

		//A status code should be returned (200 on success, 400 for invalid input")

		//Also send back the new todo as confirmation, as the spec outlines
	}
}
