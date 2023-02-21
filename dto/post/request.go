package postdto

type PostRequest struct {
	Title       string   `json:"title"  gorm:"type: varchar(255)" form:"title"`
	Description string   `json:"description" gorm:"type:text" form:"description" `
	UserID      int      `json:"user_id" form:"user_id"`
	Images      []string `json:"image" form:"images"`
}
