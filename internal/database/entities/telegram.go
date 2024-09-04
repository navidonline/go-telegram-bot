package entities

import (
	"gorm.io/gorm"
)

type DbTelegramUser struct {
	gorm.Model
	UserId       int64  `json:"id" gorm:"unique;not null"`
	RefId        int64  `json:"ref_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsBot        bool   `json:"is_bot"`
	IsPremium    bool   `json:"is_premium"`
}
