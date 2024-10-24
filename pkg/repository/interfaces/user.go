package interfaces

import (
	"github.com/Suraj18893/go-Ecommerce/pkg/domain"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
)

type UserRepository interface {
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	SignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	AddAddress(id int, address models.AddAddress, result bool) error
	GetAddresses(id int) ([]domain.Address, error)
	CheckIfFirstAddress(id int) bool
	GetUserDetails(id int) (models.UserDetailsResponse, error)
	FindUserIDByOrderID(orderID int) (int, error)
	FindUserByOrderID(orderId string)(domain.User,error)
	ChangePassword(id int, password string) error
	GetPassword(id int) (string, error)
	FindIdFromPhone(phone string) (int, error)
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error
	EditUsername(id int, username string) error

	RemoveFromCart(id int, inventoryID int) error
	ClearCart(cartID int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error

	GetCartID(id int) (int, error)
	GetProductsInCart(cart_id int) ([]int, error)
	FindProductNames(inventory_id int) (string, error)
	FindCartQuantity(cart_id, inventory_id int) (int, error)
	FindPrice(inventory_id int) (float64, error)
	FindCategory(inventory_id int) (int, error)

}
