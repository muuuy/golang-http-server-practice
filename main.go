package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	todo := Todo{ID: "1", Title: "1", Completed: false}
	todos = append(todos, todo)

	todos = append(todos, Todo{ID: "2", Title: "2", Completed: false})

	todos = append(todos, Todo{ID: "3", Title: "3", Completed: false})

	todos = append(todos, Todo{ID: "4", Title: "4", Completed: false})

	fmt.Println("The original slice is:", todos)

	// Define routes
	http.HandleFunc("/todos", getTodos)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
