package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Roll is model for sushi
type Roll struct {
	ID          string `json:"_id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

type ResponseList struct {
	Code    int    `json:"code"`
	Payload []Roll `json:"payload"`
}
type Response struct {
	Code    int  `json:"code"`
	Payload Roll `json:"payload"`
}

// init rolls var as a slice (dynamic length Array)
var rolls []Roll
var responseList ResponseList
var response Response

func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			response.Code = 200
			response.Payload = item
			json.NewEncoder(w).Encode(response)
			return
		} else {
			response.Code = 404
			json.NewEncoder(w).Encode(response)
			return
		}
	}
}

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	responseList.Code = 200
	responseList.Payload = rolls
	// fmt.Println("Learning Go on the fly")
	json.NewEncoder(w).Encode(responseList)
}

func addRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	response.Code = 201
	response.Payload = rolls[(len(rolls) - 1)]
	json.NewEncoder(w).Encode(response)
}

func main() {
	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"},
		Roll{ID: "2", ImageNumber: "9", Name: "California Roll", Ingredients: "Crab, Avocado, Cucumber, Nori, Rice"},
	)

	// router init
	router := mux.NewRouter()

	//route list
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi/add", addRoll).Methods("POST")

	//port init
	log.Fatal(http.ListenAndServe(":3001", router))
}
