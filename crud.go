package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
type Person struct {
    ID string `json:"id,omitempty"`
    Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
    City string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}
var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, _ * http.Request ) {
	_ = json.NewEncoder(w).Encode(people)
}

func GetPeopleByIdEndpoint (w http.ResponseWriter, r * http.Request ) {
	params := mux.Vars(r)
	for _, item := range people {
        if item.ID == params["id"] {
            _ = json.NewEncoder(w).Encode(item)
            return
        }
    }
	_ = json.NewEncoder(w).Encode(&Person{})
}

func CreatePeopleEndpoint (w http.ResponseWriter, r * http.Request ) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people,person)
	_ = json.NewEncoder(w).Encode(people)
}

func DeletePeopleEndpoint (w http.ResponseWriter, r * http.Request ) {
	params := mux.Vars(r)
	for index, id := range people {
		if id.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(people)
}

func UpdatePeopleEndpoint (w http.ResponseWriter, r * http.Request ) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	for index , item := range people {
		if item.ID == params ["id"]{
			person.ID = item.ID
			people[index] = person
			_ = json.NewEncoder(w).Encode(people)
			return
		}
	}
	_ = json.NewEncoder(w).Encode("Error No existe ningun usuario con este ID")
}

func main()  {
	router := mux.NewRouter()

	//Data Random
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Jose", Lastname: "Austin", Address: &Address{City: "City Y", State: "State Y"}})

	//endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}",GetPeopleByIdEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}",CreatePeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}",DeletePeopleEndpoint).Methods("DELETE")
	router.HandleFunc("/people/{id}",UpdatePeopleEndpoint).Methods("PUT")

	//Create Server and start listening and print any errors
	log.Fatal(http.ListenAndServe(":3000", router))


}
