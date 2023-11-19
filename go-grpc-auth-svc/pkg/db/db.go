package db

import (
	"log"

	"github.com/hellokvn/go-grpc-auth-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("connection db success")
	}

	db.AutoMigrate(&models.User{})

	return Handler{db}
}
