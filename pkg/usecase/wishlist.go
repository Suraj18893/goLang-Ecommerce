package usecase

import (
	"errors"

	interfaces "github.com/Suraj18893/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Suraj18893/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Suraj18893/go-Ecommerce/pkg/utils/models"
)

type wishlistUsecase struct {
	wishRepo interfaces.WishlistRepository
}

// Constructor function
func NewWishlistUsecase(wishRepo interfaces.WishlistRepository) services.WishlistUsecase {
	return &wishlistUsecase{
		wishRepo: wishRepo,
	}
}

func (wlU *wishlistUsecase) GetWishlistID(userID int) (int, error) {
	wishlistId, err := wlU.wishRepo.GetWishlistId(userID)
	if err != nil {
		return 0, err
	}
	return wishlistId, nil
}

func (wlU *wishlistUsecase) GetWishlist(id int) ([]models.GetWishlist, error) {
	// Find wishlist id
	wishlistId, err := wlU.wishRepo.GetWishlistId(id)
	if err != nil {
		return []models.GetWishlist{}, errors.New("couldn't find wishlist id from user id")
	}
	// Find products inside wishlist
	products, err := wlU.wishRepo.GetProductsInWishlist(wishlistId)
	if err != nil {
		return []models.GetWishlist{}, errors.New("couldn't find products inside wishlist")
	}
	// Find product name
	var productName []string

	for i := range products {
		prdName, err := wlU.wishRepo.FindProductNames(products[i])
		if err != nil {
			return []models.GetWishlist{}, err
		}
		productName = append(productName, prdName)
	}
	// Find price
	var productPrice []float64

	for i := range products {
		p, err := wlU.wishRepo.FindPrice(products[i])
		if err != nil {
			return []models.GetWishlist{}, err
		}
		productPrice = append(productPrice, p)
	}
	// Find category
	var category []string

	for i := range products {
		c, err := wlU.wishRepo.FindCategory(products[i])
		if err != nil {
			return []models.GetWishlist{}, errors.New("couldn't find category")
		}
		category = append(category, c)
	}

	var getWishlist []models.GetWishlist

	for i := range products {
		var get models.GetWishlist
		get.ProductName = productName[i]
		get.Price = productPrice[i]
		get.Category = category[i]

		getWishlist = append(getWishlist, get)
	}
	return getWishlist, nil
}

func (wlU *wishlistUsecase) RemoveFromWishlist(id int, inventoryID int) error {
	err := wlU.wishRepo.RemoveFromWishlist(id, inventoryID)
	if err != nil {
		return err
	}
	return nil
}

func (wlU *wishlistUsecase) AddToWishlist(user_id, inventory_id int) error {
	// Find user wishlist id
	wishlistId, err := wlU.wishRepo.GetWishlistId(user_id)
	if err != nil {
		return errors.New("couldn't find wishlist id")
	}
	// If user has no cart,create new cart
	if wishlistId == 0 {
		wishlistId, err := wlU.wishRepo.CreateNewWishlist(user_id)
		if err != nil {
			return err
		}
		// Add products to line items
		if err := wlU.wishRepo.AddWishlistItem(wishlistId, inventory_id); err != nil {
			return errors.New("add to wishlist failed")
		}

	}
	return nil
}
