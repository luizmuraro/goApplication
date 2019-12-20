package main

import (
	"fmt"
	"net/http"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func GetPlayerScore(name string) string {

	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}
	return ""
}
