package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Booking struct{
	Id int
	User string
	Members int
}

var db *gorm.DB
var err error

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to HomePage!")
    fmt.Println("Endpoint Hit: HomePage")
}


func handleRequests(){
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

func createNewBooking(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body
	w.Header().Set("Content-Type", "application/json")
    reqBody, _ := ioutil.ReadAll(r.Body)
    var booking Booking
    json.Unmarshal(reqBody, &booking)
    db.Create(&booking) 
    fmt.Println("Endpoint Hit: Creating New Booking")
    json.NewEncoder(w).Encode(booking)
}

func returnAllBookings(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	bookings := []Booking{}
	db.Find(&bookings)
	fmt.Println("Endpoint Hit: returnAllBookings")
	json.NewEncoder(w).Encode(bookings)
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	bookings := []Booking{}
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

func main() {
	DB_USER := os.Getenv("FOOTBALL_DB_USERNAME")
	DB_PASSWORD := os.Getenv("FOOTBALL_DB_PASSWORD")
	DB_NAME := os.Getenv("FOOTBALL_DB_NAME")
	dsn := fmt.Sprintf("host=localhost user=%v password=%v dbname=%v port=5432 sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(&Booking{})
    handleRequests()
}