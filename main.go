package main

import (
	"net/http"
	"todo/handlers"
	"todo/store"
)

func main() {

	// Create the Store and Task Handler
	memStore := store.NewMemStore()
	tasksHandler := handlers.NewTasksHandler(memStore)

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()
	// Register the routes and handlers
	mux.Handle("/api/tasks", tasksHandler)
	mux.Handle("/api/tasks/", tasksHandler)
	// Run the server
	http.ListenAndServe(":8080", mux)
}
