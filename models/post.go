package models

type Post struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	UserID      int                  `json:"user_id" form:"user_id"`
	User        UsersProfileResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Images      []ImageResponse      `json:"images" gorm:"foreignKey:PostID"`
	ImagesID    []int                `json:"-" form:"images_id" gorm:"-"`
}

type PostResponse struct {
	ID          int                  `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	UserID      int                  `json:"-" `
	User        UsersProfileResponse `json:"user"`
	Images      []ImageResponse      `json:"images" gorm:"foreignKey:PostID"`
	// ImagesID    []int                `json:"-" form:"images_id" gorm:"-"`

}

func (PostResponse) TableName() string {
	return "posts"
}
