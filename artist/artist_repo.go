package artist

import "Go-Mux-Smm/models"

type ArtistRepo interface {
	GetAllArtist() ([]models.Artist, error)
}
