package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Note Struct of note
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}

var noteStore = make(map[string]Note)
var id int

//GetNoteHandler Get Notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	SetAndSendResponse(w, notes, http.StatusOK)
}

//PostNoteHandler create Notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	error := json.NewDecoder(r.Body).Decode(&note)
	if error != nil {
		panic(error)
	}
	note.CreateAt = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	SetAndSendResponse(w, note, http.StatusCreated)
}

//SetAndSendResponse default send response
func SetAndSendResponse(w http.ResponseWriter, responseValue interface{}, StatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	j, error := json.Marshal(responseValue)
	if error != nil {
		panic(error)
	}
	w.WriteHeader(StatusCode)
	w.Write(j)
}

//PutNoteHandler Update Notes
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	var noteUpdate Note
	error := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if error != nil {
		panic(error)
	}
	if note, ok := noteStore[k]; ok {
		noteUpdate.CreateAt = note.CreateAt
		delete(noteStore, k)
		noteStore[k] = noteUpdate
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

//DeleteNoteHandler Delete Notes
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("No encontramos el id %s", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	Port := ":8080"

	server := &http.Server{
		Addr:           Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening in port " + Port)
	server.ListenAndServe()
}
