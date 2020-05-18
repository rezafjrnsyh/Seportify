package repositories

import (
	"Go-Mux-Smm/artist"
	"Go-Mux-Smm/models"
	"database/sql"
	"errors"
	"fmt"
)

type ArtistRepoMysqlImpl struct {
	db *sql.DB
}

func (a *ArtistRepoMysqlImpl) GetAllArtist() ([]models.Artist, error) {
	query := "SELECT * from artist"
	rows, err := a.db.Query(query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[ArtistRepoMysqlImpl.GetAllArtist] Error when query get all artist: %w", err)
	}

	defer rows.Close()

	var sliceArtist []models.Artist

	for rows.Next() {
		var each = models.Artist{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Debut, &each.Category, &each.Image)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[ArtistRepoMysqlImpl.GetAllArtist] Error when scanning rows artist: %w", err)
		}

		sliceArtist = append(sliceArtist, each)
	}

	return sliceArtist, nil
}

func CreateArtistRepoMysqlImpl(db *sql.DB) artist.ArtistRepo {
	return &ArtistRepoMysqlImpl{db}
}


