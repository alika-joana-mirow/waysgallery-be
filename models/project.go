package models

type Project struct {
	ID          int           `json:"id" gorm:"primary_key:auto_increment"`
	Description string        `json:"description"`
	File        string        `json:"file" `
	HiredId     int           `json:"hired_id"`
	Hired       HiredResponse `json:"hired"`
}

type ProjectResponse struct {
	ID          int    `json:"-"`
	Description string `json:"description"`
	File        string `json:"file" `
}

func (ProjectResponse) TableName() string {
	return "projects"
}
