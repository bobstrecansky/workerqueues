package main

import (
	"fmt"
	"net/http"
	"time"
)

type WorkRequest struct {
	Name  string
	Delay time.Duration
}

var WorkQueue = make(chan WorkRequest, 100)

func Collector(w http.ResponseWriter, r *http.Request) {

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Bad delay value: "+err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	work := WorkRequest{Name: name, Delay: delay}
	WorkQueue <- work
	fmt.Println("Work request queued")
	w.WriteHeader(http.StatusCreated)
	return
}
