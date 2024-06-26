package models

type Item struct {
	ID         int     `json:"id"`
	NameEng    string  `json:"name_eng"`
	Price      float64 `json:"price"`
	CategoryID int     `json:"category_id"`
}
