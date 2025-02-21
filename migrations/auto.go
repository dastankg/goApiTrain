package main

import (
	"apiProject/internal/link"
	"apiProject/internal/stat"
	"apiProject/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&stat.Stat{})
}
