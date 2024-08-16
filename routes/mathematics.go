package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type CalculationRequest struct {
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Operation string  `json:"operation"`
}

type CalculationResponse struct {
	Result float64 `json:"result"`
}

func Handlesum(w http.ResponseWriter, r *http.Request) {
	var req CalculationRequest

	// Decode the JSON request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Print("value gotten= ", req)

	// Perform the calculation
	var result float64
	switch req.Operation {
	case "add":
		result = req.Num1 + req.Num2
	case "subtract":
		result = req.Num1 - req.Num2
	default:
		http.Error(w, "Invalid operation", http.StatusBadRequest)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	response := CalculationResponse{Result: result}
	json.NewEncoder(w).Encode(response)
	log.Print("result sent = ", response)

}
