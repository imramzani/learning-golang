package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Roll is model for sushi
type Roll struct {
	ID          string `json:"_id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

type Response struct {
	Code    int    `json:"code"`
	Payload []Roll `json:"payload"`
}

// init rolls var as a slice (dynamic length Array)
var rolls []Roll
var response Response

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"},
		Roll{ID: "2", ImageNumber: "9", Name: "California Roll", Ingredients: "Crab, Avocado, Cucumber, Nori, Rice"},
	)
	response.Code = 200
	response.Payload = rolls
	// router init
	router := mux.NewRouter()

	//route list
	router.HandleFunc("/sushi", getRolls).Methods("GET")

	//port init
	log.Fatal(http.ListenAndServe(":3001", router))
}
