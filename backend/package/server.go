package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	ListenAddr string
}

func (s *Server) Run() {
	log.SetPrefix("Run: ")
	router := mux.NewRouter()

	router.HandleFunc("/", s.handleIndex)

	log.Println("Server is running on port", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)
}

func NewServer(addr string) *Server {
	return &Server{
		ListenAddr: addr,
	}
}

var indexPage = "<h1>Index<h1/>"

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexPage))
}

///Aux Functions
