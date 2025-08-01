package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux" // New external dependency
)

type Book struct {
	Book_Name   string
	Publication string
	Chapter     int
	Price       float64
}

type Dealer struct {
	ID       int
	Name     string
	Location string
	Sale     float64
}

var Product = []Book{
	{Book_Name: "Science", Publication: "Mittal", Chapter: 15, Price: 390.00},
	{Book_Name: "Mathematics", Publication: "Oxford", Chapter: 12, Price: 450.00},
	{Book_Name: "History", Publication: "Penguin", Chapter: 20, Price: 550.00},
	{Book_Name: "Hindi", Publication: "Mittal", Chapter: 10, Price: 350.00},
	{Book_Name: "Economy", Publication: "Oxford", Chapter: 17, Price: 300.00},
	{Book_Name: "Geography", Publication: "Mittal", Chapter: 19, Price: 250.00},
}
var Production = []Dealer{
	{ID: 101, Name: "Jack", Location: "Minnisota", Sale: 455658.5},
	{ID: 102, Name: "Luther", Location: "New York", Sale: 358858.8},
	{ID: 103, Name: "Garry", Location: "London", Sale: 625895.8},
	{ID: 104, Name: "Adrian", Location: "Paris", Sale: 9427158.5},
	{ID: 105, Name: "Mickel", Location: "Washington", Sale: 575754.4},
	{ID: 106, Name: "Virat", Location: "Delhi", Sale: 84584.86},
}

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpont hit:homepage")
	fmt.Fprintf(w, "Hey there lucky, this is the homepage!")
}

func returnAllProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProduct")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Product)
}

func returnAllProduction(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit:returnAllProduction")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Production)
}

func handleRequests() {
	// Create a new router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Register our endpoints
	log.Println("Starting server on :10000")
	myRouter.HandleFunc("/Product", returnAllProduct)
	myRouter.HandleFunc("/Production", returnAllProduction)
	myRouter.HandleFunc("/", homepage)

	// Start the server with our new router
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
