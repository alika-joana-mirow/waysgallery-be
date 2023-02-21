package models

type Image struct {
	ID     int    `json:"id"`
	Image  string `json:"image" gorm:"type:varchar(255)"`
	PostID int
	Post   PostResponse
}

type ImageResponse struct {
	ID     int    `json:"-"`
	Image  string `json:"image"`
	PostID int    `json:"-"`
}

func (ImageResponse) TableName() string {
	return "images"
}
