package interfaces

import "github.com/Suraj18893/go-Ecommerce/pkg/utils/models"

type WishlistUsecase interface {
	AddToWishlist(user_id, inventory_id int) error
	GetWishlistID(userID int) (int, error)
	GetWishlist(id int) ([]models.GetWishlist, error)
	RemoveFromWishlist(id int, inventoryID int) error
}
