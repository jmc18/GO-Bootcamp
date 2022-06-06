package models

type Pokemon struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type1      string `json:"type_1"`
	Type2      string `json:"type_2"`
	Total      string `json:"total"`
	HP         string `json:"hp"`
	Attack     string `json:"attack"`
	Defense    string `json:"defense"`
	SpAtk      string `json:"sp_atk"`
	SpDef      string `json:"sp_def"`
	Speed      string `json:"speed"`
	Generation string `json:"generation"`
	Legendary  string `json:"legendary"`
}
