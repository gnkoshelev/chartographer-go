package handlers

import (
	"internshipApplicationTemplate/internal/imageservice"
	"log"
	"net/http"
)

func CreateImage(w http.ResponseWriter, r *http.Request) {
	width := r.URL.Query().Get("width")
	log.Println(width)
	if width == "" {
		log.Println("width is empty")
		http.Error(w, "width is empty", http.StatusBadRequest)
		return
	}
	height := r.URL.Query().Get("height")
	log.Println(height)
	if height == "" {
		log.Println("height is empty")
		http.Error(w, "height is empty", http.StatusBadRequest)
		return
	}

	uuid, err := imageservice.NewImage(width, height)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(uuid))
}
