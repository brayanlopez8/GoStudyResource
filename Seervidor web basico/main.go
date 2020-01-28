package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Mundo</h1>")
}
func HolaPrueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Mundo desde prueba</h1>")
}

type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {
	msg := mensaje{
		msg: "Hola Mundo de nuevo",
	}
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", HolaMundo)
	mux.HandleFunc("/prueba", HolaPrueba)
	mux.Handle("/hola", msg)
	server := &http.Server{
		Addr:           port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening port " + port)
	server.ListenAndServe()
}
