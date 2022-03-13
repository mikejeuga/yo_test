package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ClubServer struct {
}

func NewServer() *http.Server {
	clubServer := ClubServer{}

	r := mux.NewRouter()
	r.HandleFunc("/", clubServer.home).Methods(http.MethodGet)
	return &http.Server{
		Addr: ":8095",
		Handler: r,
	}
}

func (s *ClubServer) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Tennis Club"))
}
