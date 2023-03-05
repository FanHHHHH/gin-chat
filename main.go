package main

import (
	"gin-chat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasics{})

	user := &models.UserBasics{Name: "test leofanhe"}

	// Create
	db.Create(user)

	// Read
	println(db.First(user, 1)) // find product with integer primary key

	// Update - update product's price to 200
	db.Model(user).Update(user.Password, "qwert1")
	// Update - update multiple fields

	// Delete - delete product
	// db.Delete(&product, 1)

}
