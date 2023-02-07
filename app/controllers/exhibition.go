package controllers

import (
	"fmt"
	"mvc/app/models"
)

type ExhibitionController struct {
	exhibitions []*models.Exhibition
}

func (ec *ExhibitionController) AddArtist(cid int, a *models.Artist) *models.Exhibition {
	for _, exhibition := range ec.exhibitions {
		if exhibition.Id() == cid {
			exhibition.AddArtist(a)
			return exhibition
		}
	}
	return nil
}
func (ec *ExhibitionController) GetAllExhibitions() []*models.Exhibition {
	return ec.exhibitions
}

func (ec *ExhibitionController) CreateExhibition(e *models.Exhibition) error {
	query := fmt.Sprintf("INSERT INTO exhibitions VALUES (%d, %s, %s)", e.Id, e.Name, e.StartDate)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	ec.exhibitions = append(ec.exhibitions, e)
	return nil
}
func (ec *ExhibitionController) UpdateExhibition(e *models.Exhibition) error {
	query := fmt.Sprintf("UPDATE exhibitions SET name = %s, start_date = %s WHERE id = %d", e.Name, e.StartDate, e.Id)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	for i, exhibition := range ec.exhibitions {
		if exhibition.Id() == e.Id() {
			ec.exhibitions[i] = e
		}
	}
	return nil
}
func (ec *ExhibitionController) DeleteArtist(eid int, a *models.Artist) {
	query := fmt.Sprintf("DELETE FROM exhibition WHERE eid = %s", a.DeletePaint())
	_, err := db.Exec(query)
	if err != nil {
		return
	}
	for _, exhibition := range ec.exhibitions {
		if exhibition.Id() == eid {
			if err := exhibition.DeleteArtist(a); err != nil {
				fmt.Println(err)
			}
		}
	}
}
