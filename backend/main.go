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
	todos = []Todo{} // Ensures todos is not nil to avoid error in displaying the list when empty

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

		// when a POST request is received a new Todo struct is needed to hold the incoming data, so create one
        var newTodo Todo

        // Reading the new todo sent by the frontend
        
		// r contains multiple fields. body is most commonly used to hold the main data
		// decoders in GO need to be created using the data they will decode
		decoder := json.NewDecoder(r.Body)

		// passes "newTodo" by reference, and the decoder fills it in with data reading from r.Body
		err := decoder.Decode(&newTodo)
		
		//the status code should be returned automatically by the decoder (200 on success, 400 on invalid input")

		// basic error handling, if the decoding fails, send back a bad request error
        if err != nil {
            // Send error message if decoding fails
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }

        // add the newly written Todo to the Todo array
        todos = append(todos, newTodo)

        //Also send back the new todo as confirmation, as the spec outlines
        json.NewEncoder(w).Encode(newTodo)

	//following errors with "CORS preflight" requests, I needed to include this so that initial OPTIONS requests from browsers are not rejected
	// I found this at "https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS"
	case "OPTIONS":	
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return

	//if no other cases apply
	default:
        // only GET, POST and OPTIONS methods are allowed, so refuse any other requests
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
