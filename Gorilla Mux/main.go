package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde metodo GET")
}
func PostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde metodo POST")
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde metodo PUT")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde metodo DELETE")
}

func main() {
	port := ":8080"
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", PostUser).Methods("POST")
	r.HandleFunc("/api/users", PutUser).Methods("PUT")
	r.HandleFunc("/api/users", DeleteUser).Methods("DELETE")

	server := &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening port " + port)
	server.ListenAndServe()

}
