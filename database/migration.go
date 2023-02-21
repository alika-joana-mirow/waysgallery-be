package database

import (
	"fmt"
	"waysGallery/models"
	"waysGallery/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Hired{},
		&models.Image{},
		&models.Project{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
