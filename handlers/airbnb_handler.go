package handlers

import (
	"encoding/json"
    "net/http"
    "github.com/shokuyansh/Airbnb-API/utils"
    "github.com/gorilla/mux"
    
)

type AirbnbResponse struct{
	OccupancyPercentage float64 `json:"occupancy_percentage"`
    AverageNightRate    float64 `json:"average_night_rate"`
    HighestNightRate    float64 `json:"highest_night_rate"`
    LowestNightRate     float64 `json:"lowest_night_rate"`
}

func AirbnbRoomHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    roomId := vars["roomId"]

    if roomId == "" {
        http.Error(w, "Room ID is required", http.StatusBadRequest)
        return
    }

    
    occupancy, avgRate, highRate, lowRate, err := utils.FetchAirbnbData(roomId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

  
    response := AirbnbResponse{
        OccupancyPercentage: occupancy,
        AverageNightRate:    avgRate,
        HighestNightRate:    highRate,
        LowestNightRate:     lowRate,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}