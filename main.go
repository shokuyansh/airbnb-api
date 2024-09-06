package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shokuyansh/Airbnb-API/handlers"
)
func main(){
	r := mux.NewRouter()

    r.HandleFunc("/airbnb/{roomId}", handlers.AirbnbRoomHandler).Methods("GET")

    log.Println("API is running on port 8080")
    http.ListenAndServe(":8080", r)
}