package main

import (
	"gin-chat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasics{})

	user := &models.UserBasics{}
	user.Name = "user test"
	// Create
	db.Create(&user)

	// Read
	println(db.First(&user, 1)) // find product with integer primary key

	// Update - update product's price to 200
	db.Model(&user).Update("Password", "test password")
	// Update - update multiple fields

	// Delete - delete product
	// db.Delete(&product, 1)
}
