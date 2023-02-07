package views

import (
	"encoding/json"
	"fmt"
	"mvc/app/models"
)

func PrintExhibition(e *models.Exhibition) error {
	if res, err := json.Marshal(e); err != nil {
		return err
	} else {
		fmt.Println(string(res))
		return nil
	}
}
