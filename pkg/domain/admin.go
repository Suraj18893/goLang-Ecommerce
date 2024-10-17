package domain

import "github.com/Suraj18893/go-Ecommerce/pkg/utils/models"

type Admin struct {
	ID       int    `json:"id" gorm:"unique;not null"`
	Name     string `json:"name" gorm:"validate:required"`
	Email string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}
type AdminToken struct {
	Admin        models.AdminDetailsResponse
	Token        string
	RefreshToken string
}
