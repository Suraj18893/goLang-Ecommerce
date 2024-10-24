package interfaces

import "github.com/Suraj18893/go-Ecommerce/pkg/domain"

type CategoryRepository interface{
	AddCategory(category string)(domain.Category,error)
	CheckCategory(current string)(bool,error)
	UpdateCategory(current,new string)(domain.Category,error)
	DeleteCategory(categoryId string)error
	GetCategories()([]domain.Category,error)
}