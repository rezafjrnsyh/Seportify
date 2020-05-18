package artist

import "Go-Mux-Smm/models"

type ArtistService interface {
	GetAll() ([]models.Artist, error)
}
