package services

import (
	"Go-Mux-Smm/artist"
	"Go-Mux-Smm/models"
)

type ArtistServiceImpl struct {
	artistRepo artist.ArtistRepo
}

func (m *ArtistServiceImpl) GetAll() ([]models.Artist, error) {
	return m.artistRepo.GetAllArtist()
}

func CreateArtistService(artistRepo artist.ArtistRepo) artist.ArtistService {
	return &ArtistServiceImpl{artistRepo}
}
