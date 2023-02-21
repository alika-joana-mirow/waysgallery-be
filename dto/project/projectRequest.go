package projectdto

type Project struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Description string `json:"description" form:"description"`
	File        string `json:"file" form:"file"`
	HiredID     int    `json:"hired_id" form:"hired_id"`
}
