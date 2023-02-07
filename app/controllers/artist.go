package controllers

import (
	"database/sql"
	"fmt"
	"mvc/app/models"
)

var db *sql.DB

type ArtistController struct {
	artists []*models.Artist
}

func (ac *ArtistController) GetALLArtists() []*models.Artist {
	return ac.artists
}

func (ac *ArtistController) CreateArtist(a *models.Artist) error {
	query := fmt.Sprintf("INSERT INTO artists VALUES (%s, %d, %s)", a.Name, a.Age, a.Style)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	ac.artists = append(ac.artists, a)
	return nil
}

func (ac *ArtistController) DeleteArtist(paint string) *models.Artist {
	query := fmt.Sprintf("DELETE FROM artists WHERE paint = %s")
	_, err := db.Exec(query)
	if err != nil {
		return nil
	}
	for i, artist := range ac.artists {
		if artist.DeletePaint() == paint {
			if i < len(ac.artists) {
				ac.artists = ac.artists[i:len(ac.artists)]
			}
			return artist
		}
		return artist
	}
	return nil
}
