package imagesdto

type ImagesRequest struct {
	Image []string `json:"image" form:"image" gorm:"type: varchar(255)"`
}
