package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/MeesterMarcus/go-mux/models"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

var db *gorm.DB

func HandleRequests(databaseConnection *gorm.DB) {
	db = databaseConnection
	log.Println("Starting development server at http://127.0.0.1:10000/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/new-booking", createNewBooking).Methods("POST")
	myRouter.HandleFunc("/all-bookings", returnAllBookings)
	myRouter.HandleFunc("/booking/{id}", returnSingleBooking)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to HomePage!")
    fmt.Println("Endpoint Hit: HomePage")
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
	w.Header().Set("Content-Type", "application/json")
    reqBody, _ := ioutil.ReadAll(r.Body)
    var booking models.Booking
    json.Unmarshal(reqBody, &booking)
    db.Create(&booking) 
    fmt.Println("Endpoint Hit: Creating New Booking")
    json.NewEncoder(w).Encode(booking)
}

func returnAllBookings(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	bookings := []models.Booking{}
	db.Find(&bookings)
	fmt.Println("Endpoint Hit: returnAllBookings")
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []models.Booking{}
	db.Find(&bookings)
	for _, booking := range bookings {
		// string to int
		s , err:= strconv.Atoi(key)
		if err == nil{
		   if booking.Id == s {
			  fmt.Println(booking)
			  fmt.Println("Endpoint Hit: Booking No:",key)
			  json.NewEncoder(w).Encode(booking)
		   }
		}
	 }
}