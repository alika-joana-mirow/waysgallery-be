package models

import "time"

type Hired struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	StartDate   time.Time            `json:"start_date"`
	EndDate     time.Time            `json:"end_date"`
	Price       int                  `json:"price"`
	OrderById   int                  `json:"order_byId"`
	OrderBy     UsersProfileResponse `json:"order_by"`
	OrderToId   int                  `json:"order_toId"`
	OrderTo     UsersProfileResponse `json:"order_to"`
	Status      string               `json:"status"`
}
type HiredResponse struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	StartDate   time.Time            `json:"start_date"`
	EndDate     time.Time            `json:"end_date"`
	Price       int                  `json:"price"`
	OrderById   int                  `json:"order_byId"`
	OrderBy     UsersProfileResponse `json:"order_by"`
	OrderToId   int                  `json:"order_toId"`
	OrderTo     UsersProfileResponse `json:"order_to"`
	Status      string               `json:"status"`
}

func (HiredResponse) TableName() string {
	return "hireds"
}
