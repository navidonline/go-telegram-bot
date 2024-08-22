package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TelegramUser struct {
	gorm.Model
	UserId           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsBot        bool   `json:"is_bot"`
	IsPremium    bool   `json:"is_premium"`
}

type Database struct{
	Db *gorm.DB
}

func NewDb() *Database{
	return &Database{
		Db: initDb(),
	};
}

func initDb() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&TelegramUser{})

	// Create
	//db.Create(&TelegramUser{})

	// Read
	//var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	//db.Delete(&product, 1)
	return
}
