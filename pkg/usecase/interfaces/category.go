package interfaces

import (
	"github.com/Suraj18893/go-Ecommerce/pkg/domain"
)

type CategoryUsecase interface {
	AddCategory(category string) (domain.Category, error)
	UpdateCategory(currrent, new string) (domain.Category, error)
	DeleteCategory(categoryId string) error
	GetCategories() ([]domain.Category, error)
}
