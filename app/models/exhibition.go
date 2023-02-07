package models

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

func NewExhibition(id int) *Exhibition {
	return &Exhibition{
		id:        id,
		Name:      fmt.Sprintf("exhibition %s", id),
		StartDate: time.Now(),
	}
}

type Exhibition struct {
	id        int
	Name      string    `json:"name"`
	Artists   []*Artist `json:"artists"`
	StartDate time.Time `json:"startDate"`
}

func (e Exhibition) Id() int {
	return e.id
}
func (e Exhibition) AddArtist(a *Artist) {
	query := fmt.Sprintf("INSERT INTO artists VALUES (%s, %s, %s)", a.Name, a.Age, a.Style)
	_, err := db.Exec(query)
	if err != nil {
		return
	}
	e.Artists = append(e.Artists, a)
	return
}

func (e *Exhibition) DeleteArtist(a *Artist) error {
	query := fmt.Sprintf("DELETE FROM artists WHERE name = %s", a.Name)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	for i, artist := range e.Artists {
		if artist.Name == a.Name {
			if i < len(e.Artists) {
				e.Artists = e.Artists[i:len(e.Artists)]
			}
			return nil
		}
	}
	return fmt.Errorf("artist %s not found", a.DeletePaint())
}
