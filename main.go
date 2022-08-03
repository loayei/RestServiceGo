package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var parts []Part = []Part{

}

type Part struct {
	Manufacturer string `json:"manufacturer"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	PartNumber string `json:"part_number"`
}

func createPart(w http.ResponseWriter, r *http.Request) {
	var newPart Part
	json.NewDecoder(r.Body).Decode(&newPart)

	w.Header().Set("Content-Type", "application/json")

	parts = append(parts, newPart)
	
	json.NewEncoder(w).Encode(parts)
}

func getPartList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parts)
}

func getPart(w http.ResponseWriter, r *http.Request){
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID conversion error"))
		return
		}
	
	if id >= len(parts) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	part := parts[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(part)
}

func updatePart(w http.ResponseWriter, r *http.Request){
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID conversion error"))
		return
		}
	
	if id >= len(parts) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	var updatePart Part 
	json.NewDecoder(r.Body).Decode(&updatePart)

	parts[id] = updatePart
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatePart)
}

func deletePart(w http.ResponseWriter, r *http.Request){
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID conversion error"))
		return
		}
	
	if id >= len(parts) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid id"))
		return
	}

	parts = append(parts[:id], parts[id+1:]...)
	w.WriteHeader(200)
	w.Write([]byte("Deleted"))
}

func main(){
	router:= mux.NewRouter()

	router.HandleFunc("/parts", getPartList).Methods("GET")
	router.HandleFunc("/parts", createPart).Methods("POST")
	router.HandleFunc("/parts/{id}", getPart).Methods("GET")
	router.HandleFunc("/parts/{id}", updatePart).Methods("PUT")
	router.HandleFunc("/parts/{id}", deletePart).Methods("DELETE")

	http.ListenAndServe(":6969", router)

}