package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mikejeuga/yo_test/src/domain/club"
	"github.com/mikejeuga/yo_test/src/models"
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
	r.HandleFunc("/register", clubServer.Register).Methods(http.MethodPost)
	return &http.Server{
		Addr: ":8095",
		Handler: r,
	}
}

func (s *ClubServer) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Tennis Club"))
}

func (s *ClubServer) Register(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	player := models.Player{}

	_ = json.Unmarshal(data, &player)

	_ = s.club.Add(player)

	w.WriteHeader(http.StatusCreated)
}


