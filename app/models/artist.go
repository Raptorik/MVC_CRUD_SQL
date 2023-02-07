package models

type Artist struct {
	Name         string `json:"name"`
	Painting     string `json:"paint"`
	AtExhibition bool   `json:"at_exhibition"`
	Age          int    `json:"age"`
	Style        string `json:"artist_stule"`
}

func (a *Artist) DeletePaint() string {
	return a.Painting
}
func (a *Artist) SetAtExhibition(b bool) {
	a.AtExhibition = b
}
