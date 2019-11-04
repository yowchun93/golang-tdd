package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore stores score information about players
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer is a HTTP interface
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// player := r.URL.Path[len("/players/"):]
	// fmt.Fprint(w, p.store.GetPlayerScore(player))
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
