package models

type UssdState struct {
	ID              uint `json:"id" gorm:"primary_key"`
	State           uint `json:"state"`
	Progress        uint `json:"progress"`
	MenuId          uint `json:"menu_id"`
	DifficultyLevel uint `json:"difficulty_level"`
	ConfirmFrom     uint `json:"confirm_from"`
	MenuItemId      uint `json:"menu_item_id"`
}
