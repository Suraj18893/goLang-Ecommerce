package interfaces

import (
	// "github.com/Suraj18893/go-Ecommerce/pkg/domain"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
)

type CartRepository interface {
	GetAddresses(id int) ([]models.Address, error)
	CheckIfInvAdded(invId, cartId int) bool
	GetCartId(user_id int) (int, error)
	CreateNewCart(user_id int) (int, error)
	AddLineItems(invId, cartId int) error
	AddQuantity(invId, cartId int) error
}
