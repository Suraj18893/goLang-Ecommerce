package interfaces

import (
	"github.com/Suraj18893/go-Ecommerce/pkg/domain"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
)

type UserUsecase interface {
	Login(user models.UserLogin) (models.UserToken, error)
	SignUp(user models.UserDetails) (models.UserToken, error)
	AddAddress(id int, address models.AddAddress) error
	GetAddresses(id int) ([]domain.Address, error)
	GetUserDetails(id int) (models.UserDetailsResponse, error)

	ChangePassword(id int, old string, password string, repassword string) error
	EditUser(id int, userData models.EditUser) error

	GetCartID(userID int) (int, error)
	GetCart(id int) (models.GetCartResponse, error)
	RemoveFromCart(id int, inventoryID int) error
	ClearCart(cartID int) error
	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error

	// GetWallet(id, page, limit int) (models.Wallet, error)
}
