package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"syscall"
	log "github.com/sirupsen/logrus"
	"os/signal"
)

var r *mux.Router

func Initialize() {
	r = mux.NewRouter()
}

func awaitShutdown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan
	os.Exit(0)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	res := []byte("Hello World")
	w.WriteHeader(500)
	w.Write(res)
}

func main() {
	Initialize()
    r.HandleFunc("/", HomeHandler).Methods("GET")
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	go func() {
		println("Starting server")
		if err := http.ListenAndServe(":" + port, r); err != nil {
			log.Fatal("web-server error:", err)
		}
	}()

	awaitShutdown()
}