package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pet struct {
	ID      string `json:"id,omitempty"`
	Species string `json:"species,omitempty"`
	Breed   string `json:"breed,omitempty"`
}

var pets []Pet

func GetPets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pets)
}

func GetPet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pets {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pet{})
}
func CreatePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var pet Pet
	_ = json.NewDecoder(r.Body).Decode(&pet)
	pet.ID = params["id"]
	pets = append(pets, pet)
	json.NewEncoder(w).Encode(pets)
}
func DeletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range pets {
		if item.ID == params["id"] {
			pets = append(pets[:index], pets[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pets)
	}
}

func main() {
	router := mux.NewRouter()
	pets = append(pets, Pet{ID: "1", Species: "Dog", Breed: "Chocolate Lab"})
	pets = append(pets, Pet{ID: "2", Species: "Bird", Breed: "Cockatoo"})
	pets = append(pets, Pet{ID: "3", Species: "Cat", Breed: "Siamese"})
	router.HandleFunc("/pets", GetPets).Methods("GET")
	router.HandleFunc("/pet/{id}", GetPet).Methods("GET")
	router.HandleFunc("/pet/{id}", CreatePet).Methods("POST")
	router.HandleFunc("/pet/{id", DeletePet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}
