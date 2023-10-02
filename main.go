package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var items []Item

func main() {
	// Initialize some sample data
	items = []Item{
		{ID: 1, Name: "Item 1"},
		{ID: 2, Name: "Item 2"},
		{ID: 3, Name: "Item 3"},
	}

	// Define routes
	http.HandleFunc("/items", getItems)
	http.HandleFunc("/items/add", addItem)
	http.HandleFunc("/items/update", updateItem)
	http.HandleFunc("/items/delete", deleteItem)

	// Start the server on port 8080
	fmt.Println("Server is running on :8080...")
	http.ListenAndServe(":8080", nil)
}

// Handler to get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Handler to add a new item
func addItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newItem.ID = len(items) + 1
	items = append(items, newItem)
	log.Print("hi")
	w.WriteHeader(http.StatusCreated)
}

// Handler to update an item
func updateItem(w http.ResponseWriter, r *http.Request) {
	var updatedItem Item
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.ID == updatedItem.ID {
			items[i] = updatedItem
			return
		}
	}

	http.NotFound(w, r)
}

// Handler to delete an item
func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if fmt.Sprintf("%d", item.ID) == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}
