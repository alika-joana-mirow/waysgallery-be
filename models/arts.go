package models

type Arts struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Image string `json:"image"`
}
