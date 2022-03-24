package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/yo_test/models"
	"github.com/mikejeuga/yo_test/src/internal/domain/club"
	"io"
	"net/http"
)

type ClubServer struct {
	club club.TennisClub
}

func NewServer(tennisClub club.TennisClub) *http.Server {
	clubServer := ClubServer{tennisClub}

	r := mux.NewRouter()

	r.HandleFunc("/", clubServer.home).Methods(http.MethodGet)
	r.HandleFunc("/register", clubServer.register).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":8095",
		Handler: r,
	}
}

func (s *ClubServer) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Tennis Club"))
}

func (s *ClubServer) register(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("error reading into the request bod")
	}
	defer r.Body.Close()

	player := models.Player{}

	err = json.Unmarshal(data, &player)
	if err != nil {
		fmt.Errorf("error Unmarshalling json")
	}

	err = s.club.Add(player)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	w.WriteHeader(http.StatusCreated)
}
