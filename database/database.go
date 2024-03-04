package database

import (
	"github.com/Bachry28/task-5-pbi-btpns-Muhammad-Bachry-Alhady/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=280595 dbname=btpns_golang port=5432 sslmode=disable timezone=Asia/Shanghai"
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&models.User{}, &models.Photo{})
}
