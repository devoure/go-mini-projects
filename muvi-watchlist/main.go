package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movieWatchlist struct {
	Name       string `json:"name"`
	Rating     uint   `json:"rating"`
	Id         string `json:"id"`
	DateLogged string `json:"date_logged"`
	Watched    bool   `json:"watched"`
	User       *user  `json:"logged_by"`
}

type user struct {
	UserName string `json:"username"`
}

var myWatchList []movieWatchlist

// Get all Movies Entries
func getWatchlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(myWatchList)
}

// Delete an Entry
func deleteEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range myWatchList {
		if item.Id == params["id"] {
			myWatchList = append(myWatchList[:index], myWatchList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(myWatchList)

}

// Get one Entry
func getEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range myWatchList {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

// Create a New Entry
func createEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var entry movieWatchlist
	json.NewDecoder(r.Body).Decode(&entry)
	entry.Id = strconv.Itoa(rand.Intn(10000000))
	myWatchList = append(myWatchList, entry)
	json.NewEncoder(w).Encode(myWatchList)
}

// Update an Entry
func updateWatchlist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Delete the operation
	for index, item := range myWatchList {
		if item.Id == params["id"] {
			myWatchList = append(myWatchList[:index], myWatchList[index+1:]...)
			break
		}
	}

	// Add the updated as a new Entry

	var entry movieWatchlist
	json.NewDecoder(r.Body).Decode(&entry)
	entry.Id = params["id"]
	myWatchList = append(myWatchList, entry)
	json.NewEncoder(w).Encode(myWatchList)
}

func main() {
	// Entry 1
	myWatchList = append(myWatchList, movieWatchlist{
		Name:       "Pulp Fiction",
		Rating:     5,
		Id:         "1",
		DateLogged: "2023-07-18",
		Watched:    false,
		User: &user{
			UserName: "Athumani Bakari",
		},
	})

	// Entry 2
	myWatchList = append(myWatchList, movieWatchlist{
		Name:       "Tar",
		Rating:     4,
		Id:         "2",
		DateLogged: "2023-07-18",
		Watched:    false,
		User: &user{
			UserName: "Athumani Bakari",
		},
	})

	// Entry 3
	myWatchList = append(myWatchList, movieWatchlist{
		Name:       "All Quite on The Western Front",
		Rating:     5,
		Id:         "3",
		DateLogged: "2023-07-18",
		Watched:    false,
		User: &user{
			UserName: "Athumani Bakari",
		},
	})

	// Instatiate the router
	r := mux.NewRouter()

	r.HandleFunc("/watchlist/all", getWatchlist).Methods("GET")
	r.HandleFunc("/watchlist/{id}", getEntry).Methods("GET")
	r.HandleFunc("/watchlist/add", createEntry).Methods("POST")
	r.HandleFunc("/watchlist/update/{id}", updateWatchlist).Methods("PUT")
	r.HandleFunc("/watchlist/delete/{id}", deleteEntry).Methods("DELETE")

	fmt.Printf(">>>> Starting MUVI-WATCHLIST SERVER at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
