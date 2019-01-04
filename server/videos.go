package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/komly/golangvideos_backend/service/videos"
	"log"
	"net/http"
	"strconv"
)

type getAllVideosResp struct {
	Videos []*videos.Video `json:"videos"`
}

func (s *Server) getAllVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := s.videos.GetAllVideos(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	resp := getAllVideosResp{
		Videos: videos,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) createVideo(w http.ResponseWriter, r *http.Request) {
	videoReq := &videos.Video{}
	if err := json.NewDecoder(r.Body).Decode(videoReq); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	err := s.videos.CreateVideo(r.Context(), videoReq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := json.NewEncoder(w).Encode(videoReq); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *Server) getVideoByID(w http.ResponseWriter, r *http.Request) {
	videoIDStr, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "invalid `id` param", 400)
		return
	}
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		http.Error(w, "invalid `id` param", 400)
		return
	}

	video, err := s.videos.GetVideoById(r.Context(), int64(videoID))
	if err != nil {
		log.Printf("can't get video by id: %s", err)
		http.Error(w, "internal server error", 500)
		return
	}
	if err := json.NewEncoder(w).Encode(video); err != nil {
		log.Printf("can't encode response: %s", err)
		http.Error(w, "internal server error", 500)
		return
	}
}

func (s *Server) updateVideoByID(w http.ResponseWriter, r *http.Request) {
	videoIDStr, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "invalid `id` param", 400)
		return
	}
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		http.Error(w, "invalid `id` param", 400)
		return
	}

	videoReq := &videos.Video{}
	if err := json.NewDecoder(r.Body).Decode(videoReq); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	videoReq.ID = int64(videoID)

	err = s.videos.UpdateVideoById(r.Context(), videoReq)
	if err != nil {
		if err == videos.NotFound {
			w.WriteHeader(404)
			return
		}
		log.Printf("can't update video by id: %s", err)
		http.Error(w, "internal server error", 500)
		return
	}
	if err := json.NewEncoder(w).Encode(videoReq); err != nil {
		log.Printf("can't encode response: %s", err)
		http.Error(w, "internal server error", 500)
		return
	}
}

func (s *Server) deleteVideoByID(w http.ResponseWriter, r *http.Request) {
	videoIDStr, ok := mux.Vars(r)["id"]
	if !ok {
		http.Error(w, "invalid `id` param", 400)
		return
	}
	videoID, err := strconv.Atoi(videoIDStr)
	if err != nil {
		http.Error(w, "invalid `id` param", 400)
		return
	}

	err = s.videos.DeleteVideoById(r.Context(), int64(videoID))
	if err != nil {
		if err == videos.NotFound {
			w.WriteHeader(404)
			return
		}
		log.Printf("can't update video by id: %s", err)
		http.Error(w, "internal server error", 500)
		return
	}
}
