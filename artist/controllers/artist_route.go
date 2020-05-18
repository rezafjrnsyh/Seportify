package controllers

import (
	"Go-Mux-Smm/artist"
	"Go-Mux-Smm/library"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type ArtisHandler struct {
	ArtistService artist.ArtistService
}

func CreateArtistRoutes(r *mux.Router, artistService artist.ArtistService){
	artistHandler :=ArtisHandler{artistService}

	r.HandleFunc("/", artistHandler.GetAllArtist).Methods(http.MethodGet)
}

func (m *ArtisHandler) GetAllArtist(res http.ResponseWriter, req *http.Request) {
	artist, err := m.ArtistService.GetAll()

	if err != nil {
		library.HandleError(res, "Oops, Something Went Wrong.")
		fmt.Println("[ArtistHandler.getAllArtist] error when reading request body: ", err.Error())
		return
	}

	library.HandleSuccess(res, artist)
}

