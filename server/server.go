package server

import (
	"github.com/gorilla/mux"
	"github.com/komly/golangvideos_backend/service/videos"
	"net/http"
)

type Server struct {
	videos *videos.VideosService
	mux    *mux.Router
}

func New(videos *videos.VideosService) *Server {
	mux := mux.NewRouter()
	s := &Server{
		mux:    mux,
		videos: videos,
	}
	apiMux := mux.PathPrefix("/api/v1").Subrouter()
	videosMux := apiMux.PathPrefix("/videos").Subrouter()
	videosMux.HandleFunc("/", s.getAllVideos).Methods("GET")
	videosMux.HandleFunc("/{id}", s.getVideoByID).Methods("GET")
	videosMux.HandleFunc("/{id}", s.updateVideoByID).Methods("PUT")
	videosMux.HandleFunc("/{id}", s.deleteVideoByID).Methods("DELETE")
	videosMux.HandleFunc("/", s.createVideo).Methods("POST")

	return s
}

func (s *Server) Start() error {
	if err := http.ListenAndServe(":5000", s.mux); err != nil {
		return err
	}
	return nil
}
