package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("api is on 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			Id: 1,
			Name: "Renan",
		},
		{
			Id: 2,
			Name: "Gorge",
		},
		{
			Id: 3,
			Name: "Red",
		},
		
	})
}