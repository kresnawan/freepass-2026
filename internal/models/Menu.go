package models

type Menu struct {
	MenuId    int    `json:"menu_id"`
	CanteenId int    `json:"canteen_id"`
	MenuName  string `json:"menu_name"`
	Price     int    `json:"price"`
}
