package views

import (
	"encoding/json"
	"fmt"
	"mvc/app/models"
)

func PrintArtist(a *models.Artist) error {
	if res, err := json.Marshal(a); err != nil {
		return err
	} else {
		fmt.Println(string(res))
		return nil
	}
}
